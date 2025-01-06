SELECT ride_id,
    tracking_time,
    latitude,
    longitude,
    speed,
    shakiness
FROM trackings
WHERE ride_id = $1;