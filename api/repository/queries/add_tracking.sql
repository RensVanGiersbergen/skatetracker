INSERT INTO trackings (
        ride_id,
        tracking_time,
        latitude,
        longitude,
        speed,
        shakiness
    )
VALUES ($1, $2, $3, $4, $5, $6);