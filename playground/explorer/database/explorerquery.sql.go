// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: explorerquery.sql

package database

import (
	"context"
)

const deleteBlockByID = `-- name: DeleteBlockByID :exec
DELETE FROM blocks WHERE id = ?
`

func (q *Queries) DeleteBlockByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBlockByID, id)
	return err
}

const getLatestBlock = `-- name: GetLatestBlock :one
SELECT id, height, txcount, hash FROM blocks ORDER BY height DESC LIMIT 1
`

func (q *Queries) GetLatestBlock(ctx context.Context) (Block, error) {
	row := q.db.QueryRowContext(ctx, getLatestBlock)
	var i Block
	err := row.Scan(
		&i.ID,
		&i.Height,
		&i.Txcount,
		&i.Hash,
	)
	return i, err
}

const getLimitedTransactions = `-- name: GetLimitedTransactions :many
SELECT id, cosmoshash, ethhash, typeurl, sender, blockheight FROM transactions ORDER BY id DESC LIMIT ?
`

func (q *Queries) GetLimitedTransactions(ctx context.Context, limit int64) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getLimitedTransactions, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.Cosmoshash,
			&i.Ethhash,
			&i.Typeurl,
			&i.Sender,
			&i.Blockheight,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransactions = `-- name: GetTransactions :many
SELECT id, cosmoshash, ethhash, typeurl, sender, blockheight FROM transactions
`

func (q *Queries) GetTransactions(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.Cosmoshash,
			&i.Ethhash,
			&i.Typeurl,
			&i.Sender,
			&i.Blockheight,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertBlock = `-- name: InsertBlock :one
INSERT INTO blocks(
    height,  txcount,  hash
) VALUES (
    ?, ?, ?
)
RETURNING id
`

type InsertBlockParams struct {
	Height  int64
	Txcount int64
	Hash    string
}

func (q *Queries) InsertBlock(ctx context.Context, arg InsertBlockParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertBlock, arg.Height, arg.Txcount, arg.Hash)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertTransaction = `-- name: InsertTransaction :one
INSERT INTO transactions(
    cosmoshash, ethhash, typeurl, sender, blockheight
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING id
`

type InsertTransactionParams struct {
	Cosmoshash  string
	Ethhash     string
	Typeurl     string
	Sender      string
	Blockheight int64
}

func (q *Queries) InsertTransaction(ctx context.Context, arg InsertTransactionParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertTransaction,
		arg.Cosmoshash,
		arg.Ethhash,
		arg.Typeurl,
		arg.Sender,
		arg.Blockheight,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}