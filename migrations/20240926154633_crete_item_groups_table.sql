-- +goose Up
-- +goose StatementBegin
CREATE TABLE item_groups (
     group_id SERIAL NOT NULL,
     item_id BIGINT NOT NULL,
     source_id BIGINT NOT NULL
);

ALTER TABLE item_groups
ADD PRIMARY KEY (group_id);

ALTER TABLE item_groups
ADD UNIQUE (item_id);

ALTER TABLE item_groups
ADD CONSTRAINT item_groups_source_id_foreign FOREIGN KEY (source_id) REFERENCES source_type (source_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS item_groups CASCADE;

-- +goose StatementEnd