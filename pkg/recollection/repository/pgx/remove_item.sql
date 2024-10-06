-- name: RemoveItem :exec
DELETE FROM animelayer_items
WHERE identifier = @identifier;