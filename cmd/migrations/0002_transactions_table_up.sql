CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    description TEXT NOT NULL,
    amount INT NOT NULL,
    date DATE NOT NULL
);