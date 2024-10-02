-- +goose Up
-- +goose StatementBegin
CREATE TABLE core_to_animelayer (
    id_core INT NOT NULL,
    id_animelayer INT NOT NULL
);

ALTER TABLE core_to_animelayer
ADD PRIMARY KEY (id_core, id_animelayer);

ALTER TABLE core_to_animelayer
ADD CONSTRAINT core_to_animelayer_core_item_foreign FOREIGN KEY (id_core) REFERENCES core_items (id);

ALTER TABLE core_to_animelayer
ADD CONSTRAINT core_to_animelayer_animelayer_item_foreign FOREIGN KEY (id_animelayer) REFERENCES animelayer_items (id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS core_to_animelayer CASCADE;

-- +goose StatementEnd