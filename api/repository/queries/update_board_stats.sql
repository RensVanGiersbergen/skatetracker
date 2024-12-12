UPDATE boards
SET ride_count = ride_count + 1,
    total_distance = total_distance + $2,
    top_speed = GREATEST(top_speed, $3),
    total_ridetime = total_ridetime + (make_interval(secs => $4))
WHERE board_id = $1;