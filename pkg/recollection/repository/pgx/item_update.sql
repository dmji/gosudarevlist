-- name: UpdateItem :one
UPDATE animelayer_items
SET title = coalesce(sqlc.narg('title'), title),
    release_status = coalesce(sqlc.narg('release_status'), release_status),
    last_checked_date = coalesce(
        sqlc.narg('last_checked_date'),
        last_checked_date
    ),
    created_date = coalesce(sqlc.narg('created_date'), created_date),
    updated_date = coalesce(sqlc.narg('updated_date'), updated_date),
    ref_image_cover = coalesce(sqlc.narg('ref_image_cover'), ref_image_cover),
    ref_image_preview = coalesce(
        sqlc.narg('ref_image_preview'),
        ref_image_preview
    ),
    blob_image_cover = coalesce(sqlc.narg('blob_image_cover'), blob_image_cover),
    blob_image_preview = coalesce(
        sqlc.narg('blob_image_preview'),
        blob_image_preview
    ),
    torrent_files_size = coalesce(
        sqlc.narg('torrent_files_size'),
        torrent_files_size
    ),
    notes = coalesce(sqlc.narg('notes'), notes)
WHERE animelayer_items.identifier = @identifier
RETURNING id;