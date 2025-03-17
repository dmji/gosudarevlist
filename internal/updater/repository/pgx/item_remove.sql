-- name: RemoveItem :one
DELETE FROM animelayer_items
WHERE identifier = @identifier
RETURNING id;