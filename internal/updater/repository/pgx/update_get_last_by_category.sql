-- name: GetLastCategoryUpdateItem :one
SELECT u.update_date
FROM animelayer_updates u INNER JOIN animelayer_items i ON u.item_id = i.id 
WHERE i.category = @category
ORDER BY update_date DESC
LIMIT 1;