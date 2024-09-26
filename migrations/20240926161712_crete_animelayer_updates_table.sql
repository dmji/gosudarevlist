-- +goose Up
-- +goose StatementBegin
CREATE TABLE animelayer_updates(
    update_id SERIAL NOT NULL,
    item SERIAL NOT NULL,
    date DATE NOT NULL,
    title TEXT NOT NULL,
    value_old TEXT NOT NULL,
    value_new TEXT NOT NULL
);

ALTER TABLE animelayer_updates ADD PRIMARY KEY(update_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animelayer_updates CASCADE;
-- +goose StatementEnd
