INSERT INTO rides (user_id, board_id, title, description)
VALUES ($1, $2, $3, $4)
RETURNING ride_id,
    user_id,
    board_id,
    title,
    description,
    start_time;