-- name: GetItems :many
WITH sq AS (
    SELECT unnest(
            ARRAY [@category_array::CATEGORY_ANIMELAYER[]] 
    ) as cat
), cmpl as (
     SELECT unnest(
            ARRAY [@status_array::bool[]] 
    ) as cat
)
SELECT id,
 identifier,
 title,
 case WHEN is_completed THEN true ELSE (COALESCE(updated_date, created_date) < now() - (interval '1 year'))::boolean end as is_completed,
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
WHERE category IN (SELECT cat FROM sq)
    AND case WHEN is_completed THEN true ELSE (COALESCE(updated_date, created_date) < now() - (interval '1 year'))::boolean end IN (SELECT cat FROM cmpl)
    AND (
        @search_query::text = ''
        OR SIMILARITY(title, @search_query) > 0.05
    )
ORDER BY CASE
        WHEN LENGTH(@search_query::text) > 0 THEN SIMILARITY(title, @search_query::text)
    END DESC,
    CASE
        WHEN LENGTH(@search_query::text) = 0 THEN COALESCE(updated_date, created_date)
    END DESC
LIMIT @count OFFSET @offset_count;