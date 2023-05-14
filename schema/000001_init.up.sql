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

CREATE TABLE shares (
    id SERIAL PRIMARY KEY ,
    ticker TEXT NOT NULL ,
    purchase_price FLOAT4 NOT NULL ,
    estimated_selling_price FLOAT4 NOT NULL ,
    expected_amount_of_dividends FLOAT4 NOT NULL,
    amount_of_months INTEGER NOT NULL
);

CREATE TABLE companies (
    id SERIAL PRIMARY KEY ,
    name TEXT NOT NULL ,
    dept_payments INTEGER NOT NULL ,
    depreciation INTEGER NOT NULL ,
    taxes INTEGER NOT NULL ,
    market_capitalization INTEGER NOT NULL ,
    annual_profit INTEGER NOT NULL ,
    debentures INTEGER NOT NULL ,
    revenue  INTEGER NOT NULL ,
    transaction_costs INTEGER NOT NULL ,
    available_funds INTEGER NOT NULL
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

CREATE TABLE users_shares (
    id SERIAL PRIMARY KEY ,
    user_id INTEGER NOT NULL ,
    share_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ,
    FOREIGN KEY (share_id) REFERENCES shares(id) ON DELETE CASCADE
);

CREATE TABLE users_companies (
    id SERIAL PRIMARY KEY ,
    user_id INTEGER NOT NULL ,
    company_id INTEGER NOT NULL ,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ,
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
);


