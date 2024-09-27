-- +goose Up
-- +goose StatementBegin
INSERT INTO source_type (type_name)
VALUES ('AnimeLayer');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM source_type
WHERE type_name = 'AnimeLayer';

-- +goose StatementEnd