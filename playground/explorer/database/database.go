package database

import (
	"context"
	"database/sql"
	_ "embed"

	// blank import to support sqlite3
	_ "github.com/mattn/go-sqlite3"
)

//go:embed explorerschema.sql
var ddl string

func InitExplorerDatabase(ctx context.Context, nodeDataPath string) (*sql.DB, *Queries, error) {
	db, err := sql.Open("sqlite3", nodeDataPath)
	if err != nil {
		return nil, nil, err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, nil, err
	}

	return db, New(db), nil
}