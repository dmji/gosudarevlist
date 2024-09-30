-- +goose Up
-- +goose StatementBegin
CREATE TABLE core_to_animelayer (
    id_core INT NOT NULL,
    id_animelayer INT NOT NULL
);

ALTER TABLE core_to_animelayer
ADD PRIMARY KEY (id_core, id_animelayer);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS core_to_animelayer CASCADE;

-- +goose StatementEnd