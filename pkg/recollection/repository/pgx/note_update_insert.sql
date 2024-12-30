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

-- name: InsertUpdateNoteItems :copyfrom
INSERT INTO animelayer_update_notes (update_id, title, value_old, value_new)
VALUES (@update_id, @title, @value_old, @value_new);