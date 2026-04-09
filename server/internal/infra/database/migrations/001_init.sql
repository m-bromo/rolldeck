-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);


-- +goose Up
CREATE TABLE IF NOT EXISTS characters (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    class VARCHAR(255),
    race VARCHAR(255),
    level INT DEFAULT 1,
    strengh INT,
    dexterity INT,
    constitution INT,
    wisdom INT,
    intelligence INT,
    charisma INT
);


-- +goose Down
DROP TABLE users
-- +goose Down
DROP TABLE characters
