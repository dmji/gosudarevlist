-- +goose Up
-- +goose StatementBegin
CREATE TABLE core_items (
     id SERIAL NOT NULL,
     -- optional, can't auto-search now
     mal_item_id INT NOT NULL DEFAULT -1
);

ALTER TABLE core_items
ADD PRIMARY KEY (id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS core_items CASCADE;

-- +goose StatementEnd