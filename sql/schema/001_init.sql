-- +goose Up
create table users (
    id uuid primary key,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    username text unique not null,
    password text not null,
    role text not null default 'officer' check ( role in ('admin', 'officer') ),
    clearance_level text not null
);
create table units (
    id uuid primary key,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    name text not null,
    commander_id uuid not null,
    foreign key (commander_id) references users(id)
);
create table missions(
    id uuid primary key,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    name text not null,
    objective text not null,
    start_time timestamp not null,
    end_time timestamp not null,
    outcome text not null,
    losses int not null,
    enemy_losses int not null,
    enemy_forces_size    int not null,
    own_forces_size      int not null,
    notes text,
    classification_level text not null default 'C' check ( classification_level in ('X' ,'A', 'B', 'C')),
    unit_id uuid not null,
    foreign key (unit_id) references units(id)
);

-- +goose Down
drop table missions;
drop table units;
drop table users;