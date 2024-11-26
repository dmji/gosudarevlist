-- name: GetFilters :many
WITH sq AS (
    SELECT unnest(
            ARRAY [@category_array::CATEGORY_ANIMELAYER[]] 
    ) as cat
),
cmpl as (
     SELECT unnest(
            ARRAY [@status_array::bool[]] 
    ) as cat
),
items AS (
SELECT category,
 case WHEN is_completed THEN true ELSE (COALESCE(updated_date, created_date) < now() - (interval '1 year'))::boolean end as is_completed,
 created_date,
 updated_date
FROM animelayer_items
)
SELECT DISTINCT 'category' AS "name",
    category::text AS "value",
    COUNT(category) AS "count"
FROM items
GROUP BY category
UNION
SELECT DISTINCT 'release_status' AS "name",
    CASE WHEN is_completed THEN 'completed' ELSE 'on_air' END AS "value",
    COUNT(is_completed) AS "count"
FROM items
GROUP BY is_completed;