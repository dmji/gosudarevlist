-- name: GetItemsWithSearch :many
SELECT *
FROM animelayer_items
WHERE SIMILARITY(title, @search_query) > 0.05
ORDER BY SIMILARITY(title, @search_query) DESC
LIMIT @count OFFSET @offset_count;

-- name: GetItems :many
SELECT *
FROM animelayer_items
ORDER BY GREATEST(updated_date, created_date) DESC
LIMIT @count OFFSET @offset_count;