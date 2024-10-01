-- migrate:up

create table node (
    id int primary key generated always as identity,
    label varchar not null unique
);

create table edge (
    id int primary key generated always as identity,
    src int not null,
    dst int not null,
    constraint fk_node_src foreign key (src) references node(id),
    constraint fk_node_dst foreign key (dst) references node(id)
);

-- migrate:down

drop table edge;
drop table node;
