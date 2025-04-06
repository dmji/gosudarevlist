-- +goose Up
-- +goose StatementBegin
CREATE TABLE core_to_animelayer (
    id_core INT NOT NULL,
    identifier_animelayer TEXT NOT NULL
);

ALTER TABLE core_to_animelayer
ADD PRIMARY KEY (id_core, identifier_animelayer);

ALTER TABLE core_to_animelayer
ADD CONSTRAINT core_to_animelayer_core_item_foreign FOREIGN KEY (id_core) REFERENCES core_items (id);

ALTER TABLE core_to_animelayer
ADD CONSTRAINT core_to_animelayer_animelayer_item_foreign FOREIGN KEY (identifier_animelayer) REFERENCES animelayer_items (identifier);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS core_to_animelayer CASCADE;

-- +goose StatementEnd