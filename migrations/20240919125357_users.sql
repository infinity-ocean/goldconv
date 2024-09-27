-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR
);
INSERT INTO users (id, name) VALUES 
(1, 'Grimes'),
(2, 'Merchant Schooner'),
(3, 'Matthew Bellamy'),
(4, 'Erenor Guardian'),
(5, 'Crimson Faction'),
(6, 'Dawns Light'),
(7, 'Moonlight Armoire'),
(8, 'Windswept Warden'),
(9, 'Shadowfang Talon'),
(10, 'Blazing Phoenix');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
