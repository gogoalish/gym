package repository

import (
	"database/sql"

	"project/internal/model"
)

type dbBase struct {
	db *sql.DB
}

func newDB(db *sql.DB) *dbBase {
	return &dbBase{db: db}
}

func (d *dbBase) execOne(query string, arg ...any) error {
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(arg...)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return model.ErrNotDBExecOne
	}
	return nil
}

func (d *dbBase) execMany(query string, arg ...any) error {
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(arg...)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows <= 0 {
		return model.ErrNotDBExecMany
	}
	return nil
}

func (d *dbBase) query(query string, arg ...any) (*sql.Rows, error) {
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	return stmt.Query(arg...)
}

func (d *dbBase) queryRow(query string, arg ...any) (*sql.Row, error) {
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	return stmt.QueryRow(arg...), nil
}
