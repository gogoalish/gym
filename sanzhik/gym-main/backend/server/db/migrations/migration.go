package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func CreateTable(db *sql.DB, path string) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range dir {
		info, err := file.Info()
		if err != nil {
			return err
		}
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", path, info.Name()))
		if err != nil {
			return err
		}
		if _, err := db.Exec(string(data)); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func Connect(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Create(nameDB string, path string) error {
	file, err := os.Create(fmt.Sprintf("%s.db", nameDB))
	if err != nil {
		return err
	}
	return file.Close()
}

func DropAllDB(db *sql.DB) error {
	records := `DROP TABLE IF EXISTS`

	tabls, err := SelectAllTable(db)
	if err != nil {
		return err
	}
	for _, table := range tabls {
		_, err := db.Exec(fmt.Sprintf("%s %s", records, table))
		if err != nil {
			return err
		}
	}
	return nil
}

func SelectAllTable(db *sql.DB) ([]string, error) {
	records := `SELECT name FROM sqlite_master WHERE type='table';`

	stmt, err := db.Prepare(records)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var tabls []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			return nil, err
		} else if table == "sqlite_sequence" {
			continue
		}

		tabls = append(tabls, table)
	}
	return tabls, nil
}
