-- name: InsertOrUpdateItem :exec
INSERT OR IGNORE INTO items (guid, title, category)
VALUES ('{guid}', '{title}', {category})
ON CONFLICT(guid) DO UPDATE SET title='{title}', category={category};

-- name: GetItemByGuid :one
SELECT * FROM items WHERE guid = '{guid}';