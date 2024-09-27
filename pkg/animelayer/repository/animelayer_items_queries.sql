-- name: InsertNewItem :exec
INSERT INTO animelayer_items (identifier, title, is_completed)
VALUES ($1, $2, $3);

-- name: UpdateItem :exec
UPDATE animelayer_items
SET title = $2,
    is_completed = $3
WHERE identifier = $1;

-- name: GetItemByIdentifier :one
SELECT *
FROM animelayer_items
WHERE identifier = $1;