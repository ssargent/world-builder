package repository

import (
	"context"
	"database/sql"
)

type manager struct {
}

func NewManager() Manager {
	return &manager{}
}

func (m *manager) Transaction(ctx context.Context, db WriterDB, opts *sql.TxOptions) (Transaction, error) {
	return db.BeginTxx(ctx, opts)
}
