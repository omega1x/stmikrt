package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/mailru/go-clickhouse"
	"github.com/omega1x/stmikrt/stmiklib"
)

const (
	LOGINFO = "[INFO]"
	LOGWARN = "[WARN]"
)

func main() {
	log.Println(LOGINFO, "Start the service")
	message, err := stmiklib.Get("acservice01.p12", "iOEWS3DTue")

	if err != nil {
		log.Println(LOGWARN, "Fail perform GET-method on STMIK-server. Exit code")

	} else {
		log.Println(LOGINFO, "Response message of", len(message), "byte(s) is successfully received from STMIK-server")

		log.Println(LOGINFO, "Start skimming response message")
		unit, err := stmiklib.Skim(message)
		if err != nil {
			log.Println(LOGWARN, "Fail skimming response message. Exit code")
			// return exit code
		}
		log.Println(LOGINFO, "Data for", len(unit), "automation units are found in response message")
		log.Println(LOGINFO, "Finish skimming response message")

		state, _ := stmiklib.ReadState(unit[0])
		println(len(state))

		log.Println(LOGINFO, "Start writing to Click-House Store")
		ch()
		log.Println(LOGINFO, "Finish the service. Exit code")
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

	log.Println(stmiklib.CreateQ2("stmik_messenger"))
	_, err = connect.Exec(stmiklib.CreateQ())

	if err != nil {
		log.Fatal(err)
	}

	tx, err := connect.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(stmiklib.InsertQ())

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
