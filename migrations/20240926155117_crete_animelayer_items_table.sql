-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS animelayer_items (
     item_id SERIAL NOT NULL,
     description BIGINT NULL,
     identifier TEXT NOT NULL,
     title TEXT NOT NULL,
     is_completed BOOLEAN NOT NULL,
     UNIQUE (identifier)
);

ALTER TABLE animelayer_items
ADD;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_items CASCADE;

-- +goose StatementEnd