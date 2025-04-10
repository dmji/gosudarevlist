-- name: InsertItem :exec
INSERT INTO animelayer_items (
        identifier,
        title,
        release_status,
        last_checked_date,
        first_checked_date,
        created_date,
        updated_date,
        ref_image_cover,
        ref_image_preview,
        blob_image_cover,
        blob_image_preview,
        torrent_files_size,
        notes
    )
VALUES (
        @identifier,
        @title,
        @release_status,
        @last_checked_date,
        @last_checked_date,
        @created_date,
        @updated_date,
        @ref_image_cover,
        @ref_image_preview,
        @blob_image_cover,
        @blob_image_preview,
        @torrent_files_size,
        @notes
    );