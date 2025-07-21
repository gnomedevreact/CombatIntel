-- name: CreateMission :one
insert into missions
    (id,
     name,
     objective,
     start_time,
     end_time,
     outcome,
     losses,
     enemy_losses,
     enemy_forces_size,
     own_forces_size,
     notes,
     classification_level,
     unit_id)
values (
        gen_random_uuid(),
        $1,
        $2,
        $3,
        $4,$5,
        $6,$7,
        $8,$9,
        $10, $11,
        $12
) returning *;

-- name: GetUnitMissions :many
select * from missions
where unit_id = $1;

-- name: GetMissionById :one
select * from missions
where id = $1;

-- name: DeleteMission :exec
delete from missions
where id = $1;

-- name: GetAllMissions :many
select * from missions;