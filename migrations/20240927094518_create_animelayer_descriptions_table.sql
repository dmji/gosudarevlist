-- +goose Up
-- +goose StatementBegin
CREATE TABLE animelayer_descriptions (
     description_id SERIAL NOT NULL,
     last_checked_date DATE NOT NULL,
     first_checked_date DATE NOT NULL,
     created_date TEXT NOT NULL,
     updated_date TEXT NOT NULL,
     ref_image_cover TEXT NOT NULL,
     ref_image_preview TEXT NOT NULL,
     torrent_files_size TEXT NOT NULL
);

ALTER TABLE animelayer_descriptions
ADD PRIMARY KEY (description_id);

ALTER TABLE animelayer_items
ADD CONSTRAINT animelayer_items_description_foreign FOREIGN KEY (description) REFERENCES animelayer_descriptions (description_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_descriptions CASCADE;

-- +goose StatementEnd