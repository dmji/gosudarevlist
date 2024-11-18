-- name: GetUpdates :many
SELECT animelayer_updates.item_id,
    animelayer_updates.title,
    animelayer_updates.update_date,
    animelayer_updates.value_new,
    animelayer_updates.value_old,
    animelayer_items.identifier
FROM animelayer_updates
    INNER JOIN animelayer_items ON animelayer_updates.item_id = animelayer_items.id
WHERE @show_all_categories::bool = TRUE
    OR animelayer_items.category = @category::CATEGORY_ANIMELAYER
ORDER BY animelayer_updates.update_date DESC
LIMIT @count OFFSET @offset_count;