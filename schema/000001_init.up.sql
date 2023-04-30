CREATE TABLE deposits (
    id SERIAL PRIMARY KEY ,
    initial_amount INT NOT NULL ,
    start_date DATE NOT NULL ,
    number_of_months INT NOT NULL ,
    percentage_rate REAL NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    username VARCHAR(30) NOT NULL UNIQUE ,
    email VARCHAR(50) NOT NULL UNIQUE ,
    password_hash TEXT NOT NULL
);

CREATE TABLE users_deposits (
    id SERIAL PRIMARY KEY ,
    user_id INTEGER NOT NULL ,
    deposit_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ,
    FOREIGN KEY (deposit_id) REFERENCES deposits(id) ON DELETE CASCADE
);
