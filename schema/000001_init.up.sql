CREATE TABLE deposits (
    id SERIAL PRIMARY KEY ,
    initial_amount INT NOT NULL ,
    start_date TEXT NOT NULL ,
    number_of_months INT NOT NULL ,
    percentage_rate REAL NOT NULL
);

CREATE TABLE bonds (
    id SERIAL PRIMARY KEY ,
    ticker TEXT NOT NULL ,
    amount_of_months INTEGER NOT NULL ,
    redemption_date TEXT NOT NULL ,
    size_of_coupon FLOAT4 NOT NULL ,
    number_of_payments INT2 NOT NULL ,
    purchase_price FLOAT4 NOT NULL ,
    nominal INT2 NOT NULL
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

CREATE TABLE users_bonds (
    id SERIAL PRIMARY KEY ,
    user_id INTEGER NOT NULL ,
    bond_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ,
    FOREIGN KEY (bond_id) REFERENCES bonds(id) ON DELETE CASCADE
);
