// ClickHouse Query Templates

package stmiklib

import (
	"fmt"
	"strings"
)

//CreateQ forms SQL-query for creating the table according to SIGNAL_NAMEs
func CreateQ() string {
	b := "os_id"

	return fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS example (
			country_code FixedString(2),
			%v        UInt8,
			browser_id   UInt8,
			categories   Array(Int16),
			action_day   Date,
			action_time  DateTime
		) engine=Memory
	`, b)
}

func CreateQ2(tname string) string {
	const (
		METRIC_SIGNAL_COUNT = 37
		STATUS_SIGNAL_COUNT = 65
	)
	var b strings.Builder
	for i := 0; i < METRIC_SIGNAL_COUNT+STATUS_SIGNAL_COUNT; i++ {
		if i < METRIC_SIGNAL_COUNT {
			fmt.Fprintf(&b, "\n%v INTEGER,", SIGNAL_NAME[i])
		} else {
			fmt.Fprintf(&b, "\n%v BOOL,", SIGNAL_NAME[i])
		}
	}

	return fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v (%v) engine=Memory`,
		tname, strings.TrimSuffix(b.String(), ","),
	)
}

//InsertQ forms tempate of SQL-query for inserting rows to the created table
func InsertQ() string {
	return `
			INSERT INTO example (
				country_code,
				os_id,
				browser_id,
				categories,
				action_day,
				action_time
			) VALUES (
				?, ?, ?, ?, ?, ?
			)
			`
}
