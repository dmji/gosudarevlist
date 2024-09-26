-- +goose Up
-- +goose StatementBegin
CREATE TABLE item_groups (
    group_id SERIAL NOT NULL,
    item_id SERIAL NOT NULL,
    source_type_id SERIAL NOT NULL
);

ALTER TABLE item_groups ADD PRIMARY KEY (group_id);

ALTER TABLE item_groups ADD CONSTRAINT item_groups_source_type_id_foreign FOREIGN KEY (source_type_id) REFERENCES source_type (type_id);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS item_groups CASCADE;
-- +goose StatementEnd