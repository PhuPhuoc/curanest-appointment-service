package invoicerepository

import "github.com/jmoiron/sqlx"

type invoiceRepo struct {
	db *sqlx.DB
}

func NewInvoiceRepo(db *sqlx.DB) *invoiceRepo {
	return &invoiceRepo{
		db: db,
	}
}
