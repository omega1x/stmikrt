package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/mailru/go-clickhouse"
	"github.com/omega1x/stmikrt/stmiklib"
)

const (
	LINFO = "[INFO]"
	LWARN = "[WARN]"
)

func main() {
	log.Println(LINFO, "Start the service")
	message, err := stmiklib.Get("acservice01.p12", "iOEWS3DTue")

	if err != nil {
		log.Println(LWARN, "Fail perform GET-method on STMIK-server. Exit code")

	} else {
		log.Println(LINFO, "Response message of", len(message), "byte(s) is successfully received from STMIK-server")

		/*
			f, err := os.Create("stmik.json")
			println(err)
			w := bufio.NewWriter(f)
			_, err = w.WriteString(string(message))
			w.Flush()
		*/

		//os.WriteFile("~temp.json", message, 0644)
		log.Println(LINFO, "Start parsing")
		unit, err := stmiklib.Skim(message)
		if err != nil {
			log.Println(LWARN, "Fail parse response message. Exit code")
			// return exit code
		}
		log.Println(LINFO, "Data for", len(unit), "automation units are found in response message")
		log.Println(LINFO, "Finish parse response message")

		log.Println(LINFO, "Start writing to Click-House Store")
		ch()
		log.Println(LINFO, "Finish the service. Exit code")
	}

}

func ch() {
	connect, err := sql.Open("clickhouse", "http://127.0.0.1:8123/stmik?user=default&password=default")
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS example (
			country_code FixedString(2),
			os_id        UInt8,
			browser_id   UInt8,
			categories   Array(Int16),
			action_day   Date,
			action_time  DateTime
		) engine=Memory
	`)

	if err != nil {
		log.Fatal(err)
	}

	tx, err := connect.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(`
		INSERT INTO example (
			country_code,
			os_id,
			browser_id,
			categories,
			action_day,
			action_time
		) VALUES (
			?, ?, ?, ?, ?, ?
		)`)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		if _, err := stmt.Exec(
			"RU",
			10+i,
			100+i,
			clickhouse.Array([]int16{1, 2, 3}),
			clickhouse.Date(time.Now()),
			time.Now(),
		); err != nil {
			log.Fatal(err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
