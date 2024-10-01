-- name: CreateEdge :one
insert into edge (src, dst) values ($1, $2) returning *;

-- name: GetEdge :one
select * from edge where id = $1 limit 1;

-- name: GetEdges :many
select * from edge;
