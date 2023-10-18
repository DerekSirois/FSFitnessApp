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

CREATE TABLE if not exists exercise(
  	id SERIAL PRIMARY KEY,
  	name text,
  	description text,
  	body_part_id int,            
    CONSTRAINT fk_body_part FOREIGN KEY(body_part_id)
                                    REFERENCES body_part(id)
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
