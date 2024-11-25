-- name: GetFilters :many
WITH sq AS (
    SELECT unnest(
            ARRAY [@category_array::CATEGORY_ANIMELAYER[]] 
    ) as cat
), items AS (
SELECT category,
 case WHEN is_completed THEN true ELSE (COALESCE(updated_date, created_date) < now() - (interval '1 year'))::boolean end as is_completed,
 created_date,
 updated_date
FROM animelayer_items
WHERE category IN (SELECT cat FROM sq)
    AND (
        @search_query::text = ''
        OR SIMILARITY(title, @search_query) > 0.05
    )
    AND coalesce(sqlc.narg('is_completed'), is_completed) = is_completed
ORDER BY CASE
        WHEN LENGTH(@search_query::text) > 0 THEN SIMILARITY(title, @search_query::text)
    END DESC,
    CASE
        WHEN LENGTH(@search_query::text) = 0 THEN COALESCE(updated_date, created_date)
    END DESC
)
SELECT DISTINCT 'category' AS "name",
    category::text AS "value",
    COUNT(category) AS "count"
FROM items
GROUP BY category
UNION
SELECT DISTINCT 'is_completed' AS "name",
    is_completed::text AS "value",
    COUNT(is_completed) AS "count"
FROM items
GROUP BY is_completed;