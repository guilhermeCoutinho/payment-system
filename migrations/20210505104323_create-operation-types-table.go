package main

import (
	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS operation_types (
			id INTEGER PRIMARY KEY,
			name VARCHAR
		);
		

		INSERT INTO operation_types (id, name) VALUES (0, 'cash');
		INSERT INTO operation_types (id, name) VALUES (1, 'withdraw');
		INSERT INTO operation_types (id, name) VALUES (2, 'payment');
		INSERT INTO operation_types (id, name) VALUES (3, 'installment');		
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE IF EXISTS operation_types")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20210505104323_create-operation-types-table", up, down, opts)
}
