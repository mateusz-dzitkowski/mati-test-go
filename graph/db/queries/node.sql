-- name: CreateNode :one
insert into node (label) values ($1) returning *;

-- name: GetNode :one
select * from node where id = $1 limit 1;

-- name: GetNodes :many
select * from node;
