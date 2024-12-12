INSERT INTO rides (
        user_id,
        board_id,
        completed,
        title,
        description,
        start_time,
        end_time,
        distance,
        top_speed
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING ride_id,
    user_id,
    board_id,
    completed,
    title,
    description,
    start_time,
    end_time,
    distance,
    top_speed;