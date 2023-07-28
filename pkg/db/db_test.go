package db

import (
	"fmt"
	"os"
	"testing"
)

func TestPrepareDBMustCreateTables(t *testing.T) {
	db, _ := GetDB()
	PrepareDB(&db)

	tables := []string{"user", "project"}
	for i := 0; i < len(tables); i++ {
		table := tables[i]
		row := db.QueryRow("select name from sqlite_master where name = ?", table)
		var dbTableName string
		err := row.Scan(&dbTableName)
		if err != nil {
			t.Error("happened error", err)
		}

		if dbTableName != table {
			t.Error(fmt.Sprintf("Table %s not found", table))
		}
	}

	os.Remove(databasePath)
}
