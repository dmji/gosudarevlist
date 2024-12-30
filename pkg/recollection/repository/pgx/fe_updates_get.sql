-- name: GetUpdates :many
WITH selected_categories AS (
    SELECT unnest(
            ARRAY [@category_array::CATEGORY_ANIMELAYER[]] 
    ) as cat
), selected_release_status as (
     SELECT unnest(
            ARRAY [@status_array::RELEASE_STATUS_ANIMELAYER[]] 
    ) as rs
)
SELECT animelayer_updates.id as update_id, 
    animelayer_updates.update_date as update_date,
    animelayer_updates.update_status as update_status,
    animelayer_items.identifier as item_identifier,
    animelayer_items.title as item_title
FROM animelayer_updates
    INNER JOIN animelayer_items ON animelayer_updates.item_id = animelayer_items.id
WHERE animelayer_items.category IN (SELECT cat FROM selected_categories)
    AND animelayer_items.release_status IN (SELECT rs FROM selected_release_status)
    AND (
        @search_query::text = ''
        OR SIMILARITY(animelayer_items.title, @search_query) > @similarity_threshold::float
    )
ORDER BY CASE
        WHEN LENGTH(@search_query::text) > 0 THEN SIMILARITY(animelayer_items.title, @search_query::text)
    END DESC,
    CASE
        WHEN LENGTH(@search_query::text) = 0 THEN animelayer_updates.update_date 
    END DESC
LIMIT @count::bigint OFFSET @offset_count::bigint;

-- name: GetUpdateNotes :many
SELECT title, value_old, value_new FROM animelayer_update_notes WHERE update_id = @update_id ORDER BY id;