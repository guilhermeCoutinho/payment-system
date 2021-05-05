package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID             uuid.UUID `json:"id" pg:"id, pk"`
	DocumentNumber string    `json:"document_number" pg:"document_number"`

	CreatedAt time.Time `json:"-" pg:"created_at,notnull"`
	UpdatedAt time.Time `json:"-" pg:"updated_at,notnull"`
}

type Transaction struct {
	ID            uuid.UUID     `json:"id" pg:"id, pk"`
	AccountID     uuid.UUID     `json:"account_id" pg:"account_id"`
	OperationType OperationType `json:"operation_type" pg:"operation_type"`
	Amount        int64         `json:"amount" pg:"amount"`
	Date          time.Time     `json:"date" pg:"date"`

	CreatedAt time.Time `json:"-" pg:"created_at,notnull"`
	UpdatedAt time.Time `json:"-" pg:"updated_at,notnull"`
}

type OperationType int

const (
	Cash OperationType = iota + 1
	Withdraw
	Payment
	Installment
)
