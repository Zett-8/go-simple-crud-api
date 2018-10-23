
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
    id SERIAL,
    name varchar(20),
    age int,
    created_at time,
    updated_at time,
    deleted_at time,
    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;