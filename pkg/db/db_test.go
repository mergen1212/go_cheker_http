package db

import (
	"testing"
)

func TestPrepareDBMustCreateTables(t *testing.T) {
	db, _ := GetDB()
	PrepareDB(&db)

	tables := []string{"user", "project"}
	for i := 0; i < len(tables); i++ {
		table := tables[i]
		rows, err := db.Query("select name from sqlite_master where name = ?", table)
		if err != nil {
			t.Error("Error happened", err)
		}
		defer rows.Close()

		if !rows.Next() {
			t.Error("table not found", table)
		}

		for rows.Next() {
			var dbTable string
			rows.Scan(&dbTable)
			if dbTable != table {
				t.Error("table not found", table)
			}
		}

	}
}
