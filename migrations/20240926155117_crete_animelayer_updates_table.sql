-- +goose Up
-- +goose StatementBegin
CREATE TYPE UPDATE_STATUS AS ENUM ('new', 'update', 'removed');

CREATE TABLE animelayer_updates (
     id BIGSERIAL NOT NULL,
     -- item that was updated
     item_id BIGINT NOT NULL,
     -- timestamp
     update_date timestamp NOT NULL,
     -- status
     update_status UPDATE_STATUS NOT NULL
);

ALTER TABLE animelayer_updates
ADD PRIMARY KEY (id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_updates CASCADE;

DROP TYPE UPDATE_STATUS;

-- +goose StatementEnd