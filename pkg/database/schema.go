package database

var schema = `
CREATE TABLE if not exists users(
    id SERIAL PRIMARY KEY,
    username text,
    email text,
    password bytea,
    isAdmin bool
);

CREATE TABLE  if not exists muscle(
    id SERIAL PRIMARY KEY,
    name text unique
);

CREATE TABLE if not exists exercise(
  	id SERIAL PRIMARY KEY,
  	name text,
  	description text,
  	muscle_id int,            
    CONSTRAINT fk_muscle FOREIGN KEY(muscle_id)
                                    REFERENCES muscle(id)
);

CREATE TABLE if not exists training(
  	id SERIAL PRIMARY KEY,
  	name text,
  	weekDay text,
  	user_id int,
  	CONSTRAINT fk_user FOREIGN KEY (user_id)
                                REFERENCES users(id)
);

CREATE TABLE if not exists training_exercise(
  	id SERIAL PRIMARY KEY,
  	training_id int,
  	exercise_id int,
  	CONSTRAINT fk_training FOREIGN KEY (training_id)
                                            REFERENCES training(id),
    CONSTRAINT fk_exercise FOREIGN KEY (exercise_id)
                                            REFERENCES exercise(id)
);

INSERT INTO muscle (name)
VALUES 
    ('Biceps'),
    ('Triceps'),
    ('Shoulder'),
    ('Chest'),
    ('Back'),
    ('Legs'),
    ('Abs')
on conflict (name) do nothing;
`
