-- name: InsertUpdateNote :one
INSERT INTO animelayer_updates (
        item_id,
        update_date,
        update_status
    )
VALUES (
        @item_id,
        @update_date,
        @status
    )
RETURNING id;