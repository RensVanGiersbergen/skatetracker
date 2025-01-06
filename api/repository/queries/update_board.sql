UPDATE boards
SET nickname = $2,
    brand = $3,
    primary_board = $4
WHERE board_id = $1
RETURNING board_id,
    user_id,
    nickname,
    brand,
    ride_count,
    total_distance,
    top_speed,
    CAST(
        EXTRACT(
            EPOCH
            FROM total_ridetime
        ) AS INTEGER
    ),
    primary_board,
    created_at;