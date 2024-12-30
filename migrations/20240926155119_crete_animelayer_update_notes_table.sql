-- +goose Up
-- +goose StatementBegin
CREATE TABLE animelayer_update_notes (
     id BIGSERIAL NOT NULL,
     update_id BIGINT NOT NULL,
     -- updated field
     title TEXT NOT NULL,
     value_old TEXT NOT NULL,
     value_new TEXT NOT NULL
);

ALTER TABLE animelayer_update_notes
ADD PRIMARY KEY (id);

ALTER TABLE animelayer_update_notes
ADD CONSTRAINT animelayer_updates_item_foreign FOREIGN KEY (update_id) REFERENCES animelayer_updates (id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_update_notes CASCADE;

-- +goose StatementEnd