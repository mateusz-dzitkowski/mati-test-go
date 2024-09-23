-- name: GetThing :one
select * from thing where id = $1 limit 1;

-- name: ListThings :many
select * from thing order by id;

-- name: CreateThing :one
insert into thing (created, name, amount)
values (now(), $1, $2)
returning *;
