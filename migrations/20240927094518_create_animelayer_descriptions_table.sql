-- +goose Up
-- +goose StatementBegin
CREATE TABLE animelayer_descriptions (
     description_id SERIAL NOT NULL,
     LastCheckedDate DATE NOT NULL,
     FirstCheckedDate DATE NOT NULL,
     CreatedDate TEXT NOT NULL,
     UpdatedDate TEXT NOT NULL,
     RefImageCover TEXT NOT NULL,
     RefImagePreview TEXT NOT NULL,
     TorrentFilesSize BIGINT NOT NULL
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