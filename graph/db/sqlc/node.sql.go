// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: node.sql

package db

import (
	"context"
)

const createNode = `-- name: CreateNode :one
insert into node (label) values ($1) returning id, label
`

func (q *Queries) CreateNode(ctx context.Context, label string) (Node, error) {
	row := q.db.QueryRowContext(ctx, createNode, label)
	var i Node
	err := row.Scan(&i.ID, &i.Label)
	return i, err
}

const getNode = `-- name: GetNode :one
select id, label from node where id = $1 limit 1
`

func (q *Queries) GetNode(ctx context.Context, id int32) (Node, error) {
	row := q.db.QueryRowContext(ctx, getNode, id)
	var i Node
	err := row.Scan(&i.ID, &i.Label)
	return i, err
}

const getNodes = `-- name: GetNodes :many
select id, label from node
`

func (q *Queries) GetNodes(ctx context.Context) ([]Node, error) {
	rows, err := q.db.QueryContext(ctx, getNodes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(&i.ID, &i.Label); err != nil {
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
