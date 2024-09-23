-- +goose Up
-- +goose StatementBegin
CREATE TABLE coins (
    id SERIAL PRIMARY KEY,
    balance INT,
    userFK INT REFERENCES users (id)
);
INSERT INTO coins (id, balance, userFK) VALUES (1, 583279, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE coins;
-- +goose StatementEnd
