-- name: GetAllUnits :many
select * from units;

-- name: CreateUnit :one
insert into units (id, name, commander_id)
values (
    gen_random_uuid(),
    $1,
    $2
) returning *;

-- name: AssignCommander :one
update units set commander_id = $1 returning *;

-- name: DeleteUnit :exec
delete from units where id = $1;