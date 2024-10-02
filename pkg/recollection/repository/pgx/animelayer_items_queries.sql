-- name: InsertItem :exec
INSERT INTO animelayer_items (
        identifier,
        title,
        is_completed,
        last_checked_date,
        first_checked_date,
        created_date,
        updated_date,
        ref_image_cover,
        ref_image_preview,
        torrent_files_size
    )
VALUES (
        @identifier,
        @title,
        @is_completed,
        @first_checked_date,
        @first_checked_date,
        @created_date,
        @updated_date,
        @ref_image_cover,
        @ref_image_preview,
        @torrent_files_size
    )
RETURNING id;

-- name: GetItemByIdentifier :one
SELECT *
FROM animelayer_items
WHERE identifier = @identifier;