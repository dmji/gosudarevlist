-- +goose Up
-- +goose StatementBegin
INSERT INTO source_type (source_name)
VALUES ('AnimeLayer');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM source_type
WHERE source_name = 'AnimeLayer';

-- +goose StatementEnd