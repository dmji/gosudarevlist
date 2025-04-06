-- +goose Up
-- +goose StatementBegin
CREATE TYPE RELEASE_STATUS_ANIMELAYER AS ENUM (
     'on_air',
     'incompleted',
     'completed'
);

CREATE TABLE IF NOT EXISTS animelayer_items (
     -- animelater id to form urls
     identifier TEXT NOT NULL,
     -- descriptions 
     title TEXT NOT NULL,
     release_status RELEASE_STATUS_ANIMELAYER NOT NULL,
     -- internal timestamps
     last_checked_date timestamp NOT NULL,
     first_checked_date timestamp NOT NULL,
     -- animelayer timestamps
     created_date timestamp,
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
ADD PRIMARY KEY (identifier);

ALTER TABLE animelayer_items
ADD CONSTRAINT animelayer_items_identifier_not_empty CHECK (identifier <> '');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_items CASCADE;

DROP TYPE RELEASE_STATUS_ANIMELAYER;

-- +goose StatementEnd