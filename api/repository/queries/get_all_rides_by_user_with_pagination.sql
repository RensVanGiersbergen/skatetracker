SELECT *
FROM rides
WHERE user_id = $1
ORDER BY start_time DESC
LIMIT $2 OFFSET $3;