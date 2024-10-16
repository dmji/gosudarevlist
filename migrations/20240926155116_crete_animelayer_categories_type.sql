-- +goose Up
-- +goose StatementBegin
CREATE TYPE CATEGORY_ANIMELAYER AS ENUM (
     'anime',
     'manga',
     'music',
     'dorama',
     'anime_hentai',
     'manga_hentai'
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TYPE CATEGORY_ANIMELAYER;

-- +goose StatementEnd