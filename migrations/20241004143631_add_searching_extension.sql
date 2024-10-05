-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION pg_trgm;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP EXTENSION pg_trgm;

-- +goose StatementEnd