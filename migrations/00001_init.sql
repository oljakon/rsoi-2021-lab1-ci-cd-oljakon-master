-- +goose Up

CREATE TABLE IF NOT EXISTS persons (
    id SERIAL PRIMARY KEY,
    name TEXT,
    address TEXT,
    work TEXT,
    age INT
);
