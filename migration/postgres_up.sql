CREATE TYPE user_role AS ENUM ('admin', 'user');
CREATE TYPE user_status AS ENUM ('alive', 'suspended', 'dead');
CREATE TYPE note_status AS ENUM ('active', 'archived', 'deleted');
CREATE TYPE note_category AS ENUM ('default', 'personal', 'work', 'confidential', 'temporary');
CREATE TYPE premium_plan AS ENUM ('one-month', 'three-month', 'one-year', 'free');
CREATE TYPE premium_status AS ENUM ('availed', 'cancelled');

CREATE TABLE IF NOT EXISTS user_account (
    id SERIAL,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50),
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL,
    role user_role NOT NULL,
    status user_status DEFAULT 'alive',
    PRIMARY KEY (id)
);

INSERT INTO user_account (firstname, username, email, password, role) VALUES ('admin', 'qkadmin', 'qk.admin@gmail.com', 'qkadmin', 'admin');

CREATE TABLE IF NOT EXISTS premium (
    id SERIAL,
    user_id INTEGER NOT NULL,
    plan premium_plan NOT NULL,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP NOT NULL,
    status premium_status,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user_account(id)
);

CREATE TABLE IF NOT EXISTS note (
    id SERIAL,
    user_id INTEGER NOT NULL,
    heading VARCHAR(100) NOT NULL,
    content VARCHAR(10000) NOT NULL,
    status note_status DEFAULT 'active',
    category note_category DEFAULT 'default',
    tags VARCHAR(20),
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user_account(id)
);
