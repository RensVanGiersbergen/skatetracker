SET timezone TO 'Europe/Amsterdam';
ALTER DATABASE skatetrackerdev
SET timezone = 'Europe/Amsterdam';
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(30) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);
CREATE TABLE achievements (
    achievement_id SERIAL PRIMARY KEY,
    title VARCHAR(30) UNIQUE NOT NULL,
    description VARCHAR(255),
    icon_path VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
CREATE TABLE boards (
    board_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    nickname VARCHAR(30),
    brand VARCHAR(30),
    ride_count SMALLINT DEFAULT 0,
    total_distance INTEGER DEFAULT 0,
    top_speed REAL DEFAULT 0.0,
    total_ridetime INTERVAL SECOND(0) DEFAULT '0 hours 0 minutes',
    primary_board BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
CREATE TABLE rides (
    ride_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    board_id UUID NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    title VARCHAR(30),
    description VARCHAR(255),
    start_time TIMESTAMP DEFAULT now(),
    end_time TIMESTAMP,
    distance INTEGER,
    top_speed REAL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (board_id) REFERENCES boards(board_id) ON DELETE CASCADE
);
CREATE TABLE trackings (
    tracking_id SERIAL,
    ride_id UUID NOT NULL,
    tracking_time TIMESTAMP NOT NULL DEFAULT now(),
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    speed REAL,
    shakiness REAL,
    PRIMARY KEY (tracking_id, tracking_time),
    FOREIGN KEY (ride_id) REFERENCES rides(ride_id) ON DELETE CASCADE
);
SELECT create_hypertable(
        'trackings',
        'tracking_time',
        chunk_time_interval => INTERVAL '1 day'
    );
CREATE TABLE user_achievements (
    user_id UUID NOT NULL,
    achievement_id INT NOT NULL,
    earned_at TIMESTAMP DEFAULT now(),
    PRIMARY KEY (user_id, achievement_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (achievement_id) REFERENCES achievements(achievement_id) ON DELETE CASCADE
);
INSERT INTO users (user_id, username, email, password_hash)
VALUES (
        '7f733200-d7d3-4f56-b247-ff9f1d957ef0',
        'test1',
        'test@gmail.com',
        '$2y$10$8FuXtRQH2A9pKNhgALXdCuAmUX/CYJ6ZYqS9B/xxDyfAaEOzIOGuG'
    );
INSERT INTO boards (
        board_id,
        user_id,
        nickname,
        ride_count,
        brand,
        total_distance,
        top_speed,
        total_ridetime,
        primary_board
    )
VALUES (
        'cda45722-8177-4f18-90e8-cdee6f399c1e',
        '7f733200-d7d3-4f56-b247-ff9f1d957ef0',
        'test board',
        1,
        'Verreal',
        2400,
        11.6667,
        INTERVAL '1 hours 33 minutes',
        TRUE
    );
INSERT INTO rides (
        ride_id,
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
VALUES (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '7f733200-d7d3-4f56-b247-ff9f1d957ef0',
        'cda45722-8177-4f18-90e8-cdee6f399c1e',
        TRUE,
        'test rit 1',
        'eerste test rit met test board',
        '2024-11-27 12:04:36.939431',
        '2024-11-27 13:21:01.6862090',
        2400,
        11.6667
    );
INSERT INTO trackings (
        ride_id,
        tracking_time,
        latitude,
        longitude,
        speed,
        shakiness
    )
VALUES (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:06.7559140',
        51.55973665,
        5.35546529,
        1.02,
        0.52
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:11.7022810',
        51.55975607,
        5.35554604,
        0.78,
        0.61
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:16.6726460',
        51.55976563,
        5.35552746,
        0.29,
        0.33
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:21.6624370',
        51.55988229,
        5.35566028,
        4.4,
        0.45
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:26.6643170',
        51.56003578,
        5.35591384,
        4.41,
        0.68
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:31.6902290',
        51.56008864,
        5.35605028,
        2.02,
        0.59
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:36.6752920',
        51.55971203,
        5.35614965,
        7.34,
        0.72
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:41.6844070',
        51.5590869,
        5.35614524,
        12.11,
        0.51
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:46.6811510',
        51.55853756,
        5.35627134,
        12.2,
        0.66
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:51.6810910',
        51.55803534,
        5.35673992,
        12.35,
        0.39
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:00:56.6878470',
        51.55771794,
        5.35759958,
        12.29,
        0.57
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:02.7486020',
        51.55752166,
        5.35880079,
        12.55,
        0.62
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:06.7565230',
        51.55757779,
        5.35961688,
        12.47,
        0.54
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:11.6874960',
        51.55773569,
        5.36033016,
        11.82,
        0.58
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:17.7447880',
        51.55800978,
        5.3612801,
        11.91,
        0.64
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:21.7531290',
        51.55828506,
        5.36187813,
        12.02,
        0.46
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:26.6820810',
        51.55854059,
        5.36247243,
        11.28,
        0.49
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:32.7360620',
        51.55894973,
        5.36334327,
        11.89,
        0.53
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:36.7411000',
        51.55929302,
        5.36385426,
        11.67,
        0.47
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:41.7040710',
        51.55962941,
        5.36445811,
        12.11,
        0.61
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:46.6633880',
        51.5599931,
        5.36504749,
        11.63,
        0.55
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:51.6763810',
        51.56030829,
        5.3655795,
        9.82,
        0.60
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:01:56.6725640',
        51.56057994,
        5.36604476,
        8.69,
        0.67
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:01.6783930',
        51.56090132,
        5.36669381,
        11.08,
        0.56
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:06.6874290',
        51.56114244,
        5.36748107,
        11.61,
        0.49
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:11.6804010',
        51.56140196,
        5.3681195,
        10.52,
        0.62
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:16.6861810',
        51.56164866,
        5.36878855,
        10.00,
        0.70
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:22.7419420',
        51.56198929,
        5.36968156,
        12.33,
        0.51
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:26.7393900',
        51.56226083,
        5.37021555,
        12.18,
        0.48
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:31.6864930',
        51.56261833,
        5.37069795,
        10.93,
        0.54
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:36.6920890',
        51.56300359,
        5.37104888,
        10.19,
        0.56
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:42.7474120',
        51.56376197,
        5.37125985,
        12.09,
        0.50
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:46.7462130',
        51.56418001,
        5.37159033,
        12.88,
        0.64
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:51.6955700',
        51.56459445,
        5.37206015,
        11.59,
        0.57
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:02:56.6865070',
        51.56503557,
        5.3725249,
        11.71,
        0.49
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:02.7488740',
        51.56540439,
        5.37295023,
        8.83,
        0.62
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:06.7374680',
        51.56568509,
        5.3731907,
        7.75,
        0.61
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:11.6926710',
        51.56589773,
        5.37344986,
        5.89,
        0.55
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:16.7129220',
        51.56607256,
        5.37363877,
        4.69,
        0.57
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:21.6832240',
        51.56594188,
        5.3739405,
        4.39,
        0.51
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:26.6919040',
        51.56544941,
        5.37440912,
        9.85,
        0.59
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:32.7385860',
        51.56500694,
        5.37477867,
        6.84,
        0.46
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:36.7322090',
        51.56522068,
        5.37510559,
        5.59,
        0.52
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:41.6861560',
        51.56529595,
        5.37542899,
        4.39,
        0.48
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:47.8178020',
        51.56536181,
        5.37541485,
        1.78,
        0.63
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:52.7675680',
        51.56548431,
        5.37539319,
        2.2,
        0.55
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:03:56.7700850',
        51.56545075,
        5.37532509,
        0.34,
        0.62
    ),
    (
        '40f9decb-fe66-4133-aff0-69ca962a77a1',
        '2022-04-17 13:04:01.6862090',
        51.56542665,
        5.3753347,
        0.61,
        0.50
    );