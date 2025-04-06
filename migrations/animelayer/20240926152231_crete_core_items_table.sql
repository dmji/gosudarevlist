-- +goose Up
-- +goose StatementBegin
CREATE TABLE core_items (
     -- key from mal, unique and persistent
     id BIGINT NOT NULL
     -- last update, for sorting
);

ALTER TABLE core_items
ADD PRIMARY KEY (id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS core_items CASCADE;

-- +goose StatementEnd