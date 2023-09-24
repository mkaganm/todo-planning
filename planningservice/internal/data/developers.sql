CREATE TABLE developers (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) NOT NULL,
                            duration INTERVAL NOT NULL,
                            difficulty VARCHAR(10) NOT NULL,
                            estimation INTEGER
);

INSERT INTO developers (name, duration, difficulty)
VALUES
    ('DEV1', '1 hour', '1x'),
    ('DEV2', '1 hour', '2x'),
    ('DEV3', '1 hour', '3x'),
    ('DEV4', '1 hour', '4x'),
    ('DEV5', '1 hour', '5x');

UPDATE developers SET
    estimation = 45 * EXTRACT(HOUR FROM duration) * CAST(LEFT(difficulty, LENGTH(difficulty) - 1) AS INTEGER);