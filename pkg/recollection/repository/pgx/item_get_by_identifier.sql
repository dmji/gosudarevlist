-- name: GetItemByIdentifier :one
SELECT *
FROM animelayer_items
WHERE identifier = @identifier;