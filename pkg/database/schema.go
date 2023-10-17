package database

var schema = `
CREATE TABLE if not exists users(
    id SERIAL PRIMARY KEY,
    username text,
    email text,
    password bytea
);
`
