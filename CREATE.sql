CREATE TABLE IF NOT EXISTS logging (
    id serial PRIMARY KEY,
    date TIMESTAMP UNIQUE NOT NULL,
    url VARCHAR (255) NOT NULL,
    method VARCHAR (4) NOT NULL,
    status INT NOT NULL,
    user_id serial,
    body VARCHAR (255),
    comment VARCHAR (255)
);