
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE todos (
    id SERIAL,
    name varchar(10),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    primary key(id)
);

INSERT INTO todos (id, name, created_at, updated_at, deleted_at) VALUES (
    1,
    'cook',
    to_timestamp('2006-4-27','yyyy/mm/dd'),
    to_timestamp('2006-4-27','yyyy/mm/dd'),
    null
),(
    2,
    'buy milk',
    to_timestamp('2006-5-27','yyyy/mm/dd'),
    to_timestamp('2006-5-27','yyyy/mm/dd'),
    null
),(
    3,
    'clean room',
    to_timestamp('2006-6-27','yyyy/mm/dd'),
    to_timestamp('2006-6-27','yyyy/mm/dd'),
    null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE todos;