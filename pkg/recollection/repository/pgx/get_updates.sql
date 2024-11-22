-- name: GetUpdates :many
SELECT animelayer_updates.update_date,
    animelayer_updates.update_status,
    animelayer_items.identifier,
    animelayer_items.title
FROM animelayer_updates
    INNER JOIN animelayer_items ON animelayer_updates.item_id = animelayer_items.id
WHERE @show_all_categories::bool = TRUE
    OR animelayer_items.category = @category::CATEGORY_ANIMELAYER
ORDER BY animelayer_updates.update_date DESC
LIMIT @count OFFSET @offset_count;