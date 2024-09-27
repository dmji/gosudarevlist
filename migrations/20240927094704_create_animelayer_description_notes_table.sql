-- +goose Up
-- +goose StatementBegin
CREATE TABLE animelayer_description_notes (
     note_id SERIAL NOT NULL,
     description_id BIGINT NOT NULL,
     field_name TEXT NOT NULL,
     field_text TEXT NOT NULL
);

ALTER TABLE animelayer_description_notes
ADD PRIMARY KEY (note_id);

ALTER TABLE animelayer_description_notes
ADD CONSTRAINT animelayer_description_notes_description_foreign FOREIGN KEY (description_id) REFERENCES animelayer_descriptions (description_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_description_notes CASCADE;

-- +goose StatementEnd