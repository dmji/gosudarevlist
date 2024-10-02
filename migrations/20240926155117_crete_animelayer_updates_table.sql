-- +goose Up
-- +goose StatementBegin
CREATE TABLE animelayer_updates (
     id SERIAL NOT NULL,
     -- item that was updated
     item_id INT NOT NULL,
     -- timestamp
     update_date DATE NOT NULL,
     -- updated field
     title TEXT NOT NULL,
     value_old TEXT NOT NULL,
     value_new TEXT NOT NULL
);

ALTER TABLE animelayer_updates
ADD PRIMARY KEY (id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_updates CASCADE;

-- +goose StatementEnd