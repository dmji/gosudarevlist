-- +goose Up
-- +goose StatementBegin
CREATE TABLE source_type (
     type_id SERIAL NOT NULL,
     type_name CHAR(255) NOT NULL
);

ALTER TABLE source_type
ADD PRIMARY KEY (type_id);

ALTER TABLE source_type
ADD UNIQUE (type_name);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS source_type CASCADE;

-- +goose StatementEnd