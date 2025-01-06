-- name: GetFilters :many
WITH sq AS (
    SELECT unnest(
            ARRAY [@category_array::CATEGORY_ANIMELAYER[]] 
    ) as s
),
cmpl AS (
     SELECT unnest(
            ARRAY [@status_array::RELEASE_STATUS_ANIMELAYER[]] 
    ) as c
),
items AS (
    SELECT category,
    release_status,
    created_date,
    updated_date
    FROM animelayer_items
)
SELECT DISTINCT 'category' AS "name",
    category::text AS "value",
    COUNT(category) AS "count",
    category in (SELECT s FROM sq) as "selected"
FROM items
GROUP BY category
UNION
SELECT DISTINCT 'release_status' AS "name",
    release_status::text as "value",
    COUNT(release_status) AS "count",
    release_status in (SELECT c FROM cmpl) as "selected"
FROM items
GROUP BY release_status;