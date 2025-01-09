-- name: GetItems :many
WITH selected_categories AS (
    SELECT unnest(
            ARRAY [@category_array::CATEGORY_ANIMELAYER[]] 
    ) as cat
), selected_release_status as (
     SELECT unnest(
            ARRAY [@status_array::RELEASE_STATUS_ANIMELAYER[]] 
    ) as rs
)
SELECT id,
 identifier,
 title,
 release_status,
 last_checked_date,
 first_checked_date,
 created_date,
 updated_date,
 ref_image_cover,
 ref_image_preview,
 blob_image_cover,
 blob_image_preview,
 torrent_files_size,
 notes,
 category
FROM animelayer_items
WHERE category IN (SELECT cat FROM selected_categories)
    AND release_status IN (SELECT rs FROM selected_release_status)
    AND (
        @search_query::text = ''
        OR SIMILARITY(title, @search_query) > @similarity_threshold::float
    )
ORDER BY CASE
        WHEN LENGTH(@search_query::text) > 0 THEN SIMILARITY(title, @search_query::text)
    END DESC,
    CASE
        WHEN LENGTH(@search_query::text) = 0 THEN COALESCE(updated_date, created_date)
    END DESC
LIMIT @count::bigint OFFSET @offset_count::bigint;