CREATE TABLE user_account (
id INTEGER,
username VARCHAR(50) NOT NULL,
email VARCHAR(50) NOT NULL UNIQUE,
password VARCHAR(10) NOT NULL,
status INTEGER DEFAULT 1,
PRIMARY KEY (id)
);

CREATE TABLE note (
    id SERIAL,
    user_id INTEGER NOT NULL,
    title VARCHAR(50) NOT NULL,
    text VARCHAR(1000) NOT NULL,
    status INTEGER DEFAULT 1,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user_account(id)
);
