-- name: GetUpdates :many
WITH sq AS (
    SELECT unnest(
            ARRAY [@category_array::CATEGORY_ANIMELAYER[]] 
    ) as cat
), cmpl as (
     SELECT unnest(
            ARRAY [@status_array::bool[]] 
    ) as cat
)
SELECT animelayer_updates.update_date,
    animelayer_updates.update_status,
    animelayer_items.identifier,
    animelayer_items.title
FROM animelayer_updates
    INNER JOIN animelayer_items ON animelayer_updates.item_id = animelayer_items.id
WHERE animelayer_items.category IN (SELECT cat FROM sq)
    AND case WHEN animelayer_items.is_completed THEN true ELSE (COALESCE(animelayer_items.updated_date, animelayer_items.created_date) < now() - (interval '1 year'))::boolean end IN (SELECT cat FROM cmpl)
    AND (
        @search_query::text = ''
        OR SIMILARITY(animelayer_items.title, @search_query) > 0.05
    )
ORDER BY CASE
        WHEN LENGTH(@search_query::text) > 0 THEN SIMILARITY(animelayer_items.title, @search_query::text)
    END DESC,
    CASE
        WHEN LENGTH(@search_query::text) = 0 THEN animelayer_updates.update_date 
    END DESC
LIMIT @count::bigint OFFSET @offset_count::bigint;