-- name: InsertUpdateNote :exec
INSERT INTO animelayer_updates (
        item_id,
        update_date,
        title,
        value_old,
        value_new
    )
VALUES (
        @item_id,
        @update_date,
        @update_title,
        @value_old,
        @value_new
    );

-- name: GetUpdateNote :many
SELECT *
FROM animelayer_updates
ORDER BY update_date DESC
LIMIT @count::bigint OFFSET @offset_count::bigint;