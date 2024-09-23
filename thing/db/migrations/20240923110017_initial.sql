-- migrate:up

create table thing (
    id int primary key generated always as identity,
    created timestamp with time zone not null,
    name varchar not null,
    amount int not null
)

-- migrate:down

drop table thing;
