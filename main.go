//go:generate fileb0x filebox.json
package main

import (
	"os"
	fp "path/filepath"

	"github.com/RadhiFadlillah/shiori/cmd"
	db "github.com/RadhiFadlillah/shiori/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	databasePath := "shiori.db"
	if value, found := os.LookupEnv("ENV_SHIORI_DB"); found {
		// If ENV_SHIORI_DB is directory, append "shiori.db" as filename
		if f1, err := os.Stat(value); err == nil && f1.IsDir() {
			value = fp.Join(value, "shiori.db")
		}

		databasePath = value
	}

	sqliteDB, err := db.OpenSQLiteDatabase(databasePath)
	checkError(err)

	cmd.DB = sqliteDB
	cmd.Execute()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
