// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: container.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const addContainer = `-- name: AddContainer :one
INSERT INTO container(name, password)
VALUES ($1, $2)
RETURNING container_uuid
`

type AddContainerParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (q *Queries) AddContainer(ctx context.Context, arg AddContainerParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, addContainer, arg.Name, arg.Password)
	var container_uuid uuid.UUID
	err := row.Scan(&container_uuid)
	return container_uuid, err
}

const deleteContainer = `-- name: DeleteContainer :exec
DELETE FROM container
WHERE container_uuid = $1
`

func (q *Queries) DeleteContainer(ctx context.Context, containerUuid uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteContainer, containerUuid)
	return err
}

const updateContainerName = `-- name: UpdateContainerName :one
UPDATE container
SET name = $1
WHERE container_uuid = $2
RETURNING container_uuid
`

type UpdateContainerNameParams struct {
	Name          string    `json:"name"`
	ContainerUuid uuid.UUID `json:"container_uuid"`
}

func (q *Queries) UpdateContainerName(ctx context.Context, arg UpdateContainerNameParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, updateContainerName, arg.Name, arg.ContainerUuid)
	var container_uuid uuid.UUID
	err := row.Scan(&container_uuid)
	return container_uuid, err
}

const updateContainerPassword = `-- name: UpdateContainerPassword :one
UPDATE container
SET password = $1
WHERE container_uuid = $2
RETURNING container_uuid
`

type UpdateContainerPasswordParams struct {
	Password      string    `json:"password"`
	ContainerUuid uuid.UUID `json:"container_uuid"`
}

func (q *Queries) UpdateContainerPassword(ctx context.Context, arg UpdateContainerPasswordParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, updateContainerPassword, arg.Password, arg.ContainerUuid)
	var container_uuid uuid.UUID
	err := row.Scan(&container_uuid)
	return container_uuid, err
}
