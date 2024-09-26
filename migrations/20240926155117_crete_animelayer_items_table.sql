-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS animelayer_items (
    item_id SERIAL NOT NULL,
    description bigint NULL,
    identifier TEXT NOT NULL,
    title TEXT NOT NULL,
    is_completed BOOLEAN NOT NULL
);

ALTER TABLE animelayer_items ADD UNIQUE (identifier);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_items CASCADE;
-- +goose StatementEnd