INSERT INTO Boards (user_id, nickname, brand, primary_board)
VALUES ($1, $2, $3, $4)
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