package main

import (
	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
		ALTER TABLE accounts
		ADD COLUMN credit_limit INTEGER;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE accounts DROP COLUMN credit_limit;")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20210518220832_add-column-credit-limit", up, down, opts)
}
