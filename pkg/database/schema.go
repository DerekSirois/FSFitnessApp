package database

var schema = `
CREATE TABLE if not exists users(
    id SERIAL PRIMARY KEY,
    username text,
    email text,
    password bytea,
    isAdmin bool
);

CREATE TABLE  if not exists body_part(
    id SERIAL PRIMARY KEY,
    name text unique
);

INSERT INTO body_part (name)
VALUES 
    ('Arm'),
    ('Shoulder'),
    ('Chest'),
    ('Back'),
    ('Legs'),
    ('Abs')
on conflict (name) do nothing;
`
