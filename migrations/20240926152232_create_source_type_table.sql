-- +goose Up
-- +goose StatementBegin
CREATE TABLE source_type (
     source_id SERIAL NOT NULL,
     source_name CHAR(255) NOT NULL
);

ALTER TABLE source_type
ADD PRIMARY KEY (source_id);

ALTER TABLE source_type
ADD UNIQUE (source_name);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS source_type CASCADE;

-- +goose StatementEnd