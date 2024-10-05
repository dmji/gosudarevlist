-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS animelayer_items (
     id SERIAL NOT NULL,
     -- animelater id to form urls
     identifier TEXT NOT NULL,
     -- descriptions 
     title TEXT NOT NULL,
     is_completed BOOLEAN NOT NULL,
     -- internal timestamps
     last_checked_date timestamp NOT NULL,
     first_checked_date timestamp NOT NULL,
     -- animelayer timestamps
     created_date timestamp NOT NULL,
     updated_date timestamp,
     -- static urls to files
     ref_image_cover TEXT NOT NULL,
     ref_image_preview TEXT NOT NULL,
     -- blob identificator for internal files replication
     blob_image_cover TEXT NOT NULL,
     blob_image_preview TEXT NOT NULL,
     -- torrent meta data
     torrent_files_size TEXT NOT NULL,
     -- notes
     notes TEXT NOT NULL
);

ALTER TABLE animelayer_items
ADD PRIMARY KEY (id);

ALTER TABLE animelayer_items
ADD UNIQUE (identifier);

ALTER TABLE animelayer_updates
ADD CONSTRAINT animelayer_updates_item_foreign FOREIGN KEY (item_id) REFERENCES animelayer_items (id);

ALTER TABLE animelayer_items
ADD CONSTRAINT animelayer_items_identifier_not_empty CHECK (identifier <> '');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_items CASCADE;

-- +goose StatementEnd