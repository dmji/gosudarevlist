-- name: GetUpdates :many
WITH sq AS (
    SELECT unnest(
            ARRAY [@category_array::CATEGORY_ANIMELAYER[]] 
    ) as cat
)
SELECT animelayer_updates.update_date,
    animelayer_updates.update_status,
    animelayer_items.identifier,
    animelayer_items.title
FROM animelayer_updates
    INNER JOIN animelayer_items ON animelayer_updates.item_id = animelayer_items.id
WHERE animelayer_items.category IN (SELECT cat FROM sq)
    AND (
        @search_query::text = ''
        OR SIMILARITY(animelayer_items.title, @search_query) > 0.05
    )
    AND coalesce(sqlc.narg('is_completed'), animelayer_items.is_completed) = animelayer_items.is_completed
ORDER BY CASE
        WHEN LENGTH(@search_query::text) > 0 THEN SIMILARITY(animelayer_items.title, @search_query::text)
    END DESC,
    CASE
        WHEN LENGTH(@search_query::text) = 0 THEN animelayer_updates.update_date 
    END DESC
LIMIT @count OFFSET @offset_count;