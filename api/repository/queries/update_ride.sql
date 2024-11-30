UPDATE rides
SET completed = $2,
    title = $3,
    description = $4,
    start_time = $5,
    end_time = $6,
    distance = $7,
    top_speed = $8
WHERE ride_id = $1
RETURNING *;