-- +goose Up
-- +goose StatementBegin
CREATE TABLE animelayer_notes (
     note_id SERIAL NOT NULL,
     -- item that the note applies to
     item_id INT NOT NULL,
     -- field
     field_name TEXT NOT NULL,
     field_text TEXT NOT NULL
);

ALTER TABLE animelayer_notes
ADD PRIMARY KEY (note_id);

ALTER TABLE animelayer_notes
ADD CONSTRAINT animelayer_notes_item_foreign FOREIGN KEY (item_id) REFERENCES animelayer_items (id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_notes CASCADE;

-- +goose StatementEnd