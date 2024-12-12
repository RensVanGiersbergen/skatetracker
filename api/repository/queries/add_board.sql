INSERT INTO Boards (user_id, nickname, brand)
VALUES ($1, $2, $3)
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