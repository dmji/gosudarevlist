-- name: InsertNewDescription :exec
INSERT INTO animelayer_descriptions (
        last_checked_date,
        first_checked_date,
        created_date,
        updated_date,
        ref_image_cover,
        ref_image_preview,
        torrent_files_size
    )
VALUES (
        @first_checked_date,
        @first_checked_date,
        @created_date,
        @updated_date,
        @ref_image_cover,
        @ref_image_preview,
        @torrent_files_size
    );

-- name: UpdateDescription :exec
UPDATE animelayer_descriptions
SET last_checked_date = @last_checked_date,
    created_date = @created_date,
    updated_date = @updated_date,
    ref_image_cover = @ref_image_cover,
    ref_image_preview = @ref_image_preview,
    torrent_files_size = @torrent_files_size
WHERE description_id = @description_id;

-- name: GetDescriptionByIdentifier :one
SELECT *
FROM animelayer_items
WHERE identifier = @identifier;