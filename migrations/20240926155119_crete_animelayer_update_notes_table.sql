-- +goose Up
-- +goose StatementBegin
CREATE TABLE animelayer_update_notes (
     update_id BIGINT NOT NULL,
     -- updated field
     title TEXT NOT NULL,
     value_old TEXT NOT NULL,
     value_new TEXT NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_update_notes CASCADE;

-- +goose StatementEnd