package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Connect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Create(nameDB string, path string) error {
	file, err := os.Create(fmt.Sprintf("%s.sqlite", nameDB))
	if err != nil {
		return err
	}
	return file.Close()
}

func DownDB(db *sql.DB, path string) error {
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
