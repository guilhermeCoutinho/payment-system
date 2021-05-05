package main

import (
	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS accounts (
				id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
				document_number VARCHAR NOT NULL,

				created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
			)
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE IF EXISTS accounts")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20210505103953_create-accounts-table", up, down, opts)
}
