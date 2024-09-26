-- +goose Up
-- +goose StatementBegin
CREATE TABLE coins (
    id SERIAL PRIMARY KEY,
    balance INT,
    userFK INT REFERENCES users (id)
);
INSERT INTO coins (id, balance, userFK) VALUES 
(1, 583279, 1),
(2, 320150, 2),
(3, 999998565, 3), 
(4, 912450, 4),
(5, 540246, 5),
(6, 125500, 6),
(7, 334740, 7),
(8, 455000, 8),
(9, 78624, 9), 
(10, 999999, 10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE coins;
-- +goose StatementEnd
