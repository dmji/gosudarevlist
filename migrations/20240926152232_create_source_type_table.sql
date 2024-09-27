-- +goose Up
-- +goose StatementBegin
CREATE TABLE source_type (
     type_id SERIAL NOT NULL,
     type_name CHAR(255) NOT NULL,
     PRIMARY KEY (type_id),
     UNIQUE (type_name)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS source_type CASCADE;

-- +goose StatementEnd