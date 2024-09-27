-- name: InsertNewItem :exec
INSERT INTO animelayer_items (identifier, title, is_completed)
VALUES (@identifier, @title, @is_completed);

-- name: UpdateItem :exec
UPDATE animelayer_items
SET title = @title,
    is_completed = @is_completed
WHERE identifier = @identifier;

-- name: GetItemByIdentifier :one
SELECT *
FROM animelayer_items
WHERE identifier = @identifier;