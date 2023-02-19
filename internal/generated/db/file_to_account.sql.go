// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: file_to_account.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const fileToAccount = `-- name: FileToAccount :exec
INSERT INTO file_to_account(file_uuid, account_id) 
VALUES($1, $2)
`

type FileToAccountParams struct {
	FileUuid  string    `json:"file_uuid"`
	AccountID uuid.UUID `json:"account_id"`
}

func (q *Queries) FileToAccount(ctx context.Context, arg FileToAccountParams) error {
	_, err := q.db.Exec(ctx, fileToAccount, arg.FileUuid, arg.AccountID)
	return err
}