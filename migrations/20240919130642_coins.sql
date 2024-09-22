-- +goose Up
-- +goose StatementBegin
CREATE TABLE coins (
    id SERIAL PRIMARY KEY,
    balance INT,
    userFK INT REFERENCES users (id) -- TODO: to user_id
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE coins;
-- +goose StatementEnd
