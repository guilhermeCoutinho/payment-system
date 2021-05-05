package main

import (
	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			account_id UUID NOT NULL REFERENCES accounts (id) ON DELETE CASCADE,
			operation_type INTEGER NOT NULL REFERENCES operation_types (id),
			amount INTEGER NOT NULL,
			date TIMESTAMP WITH TIME ZONE DEFAULT now(),

			created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
		)`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE IF EXISTS transactions")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20210505104623_create-transactions-table", up, down, opts)
}
