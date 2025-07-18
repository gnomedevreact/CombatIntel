-- name: GetUsers :many
select * from users;

-- name: CreateUser :one
insert into users (id, password, clearance_level, username)
values (
        gen_random_uuid(),
        $1,
        $2,
        $3
) returning *;

-- name: GetUserByUsername :one
select * from users
where username = $1;

-- name: GetUserById :one
select * from users
where id = $1;