-- +goose Up
-- +goose StatementBegin
CREATE TYPE UPDATE_STATUS AS ENUM ('new', 'update', 'removed');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TYPE UPDATE_STATUS;

-- +goose StatementEnd