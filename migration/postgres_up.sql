CREATE TYPE user_role AS ENUM ('admin', 'user');

CREATE TABLE IF NOT EXISTS user_account (
    id SERIAL,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(10) NOT NULL,
    role user_role NOT NULL,
    status INTEGER DEFAULT 1,
    PRIMARY KEY (id)
);

INSERT INTO user_account (username, email, password, role) VALUES ('qkadmin', 'qk.admin@gmail.com', 'qkadmin', 'admin');

CREATE TABLE IF NOT EXISTS note (
    id SERIAL,
    user_id INTEGER NOT NULL,
    title VARCHAR(50) NOT NULL,
    text VARCHAR(1000) NOT NULL,
    status INTEGER DEFAULT 1,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user_account(id)
);
