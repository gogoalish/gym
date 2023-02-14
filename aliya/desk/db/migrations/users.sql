CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    description VARCHAR(225),
    email VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    location VARCHAR(64),
    phone VARCHAR(64)
);
