-- name: InsertNote :exec
INSERT INTO animelayer_notes (item_id, field_name, field_text)
VALUES (
        @item_id,
        @field_name,
        @field_text
    );

-- name: UpdateNote :exec
UPDATE animelayer_notes
SET field_text = @field_text
WHERE item_id = @item_id
    AND field_name = @field_name;

-- name: GetNote :many
SELECT (field_name, field_text)
FROM animelayer_notes
WHERE item_id = @item_id;

-- name: DeleteNote :exec
DELETE FROM animelayer_notes
WHERE field_name = @field_name
    AND item_id = @item_id;