-- name: GetFilters :many
WITH checked_categories AS (
    SELECT unnest(
            ARRAY [@checked_category_array::CATEGORY_ANIMELAYER[]] 
    ) as cat
),
checked_release_status AS (
     SELECT unnest(
            ARRAY [@checked_status_array::RELEASE_STATUS_ANIMELAYER[]] 
    ) as rs
),
-- Collect all items
items AS (
    SELECT id,
        category,
        release_status,
        created_date,
        updated_date
    FROM animelayer_items
),
-- Collect filtered items
selected_categories AS (
    SELECT unnest(
            ARRAY [@selected_category_array::CATEGORY_ANIMELAYER[]] 
    ) as cat
), 
selected_release_status as (
     SELECT unnest(
            ARRAY [@selected_status_array::RELEASE_STATUS_ANIMELAYER[]] 
    ) as rs
),
items_filtered AS (
    SELECT id,
        category,
        release_status,
        created_date,
        updated_date
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
),
-- Collect filters table
categories_filtered AS (
    SELECT DISTINCT 'category' AS "name",
        category::text AS "value",
        COUNT(category) AS "count"
    FROM items_filtered
    GROUP BY category
),
categories AS (
    SELECT DISTINCT 'category' AS "name",
        category::text AS "value",
        COUNT(category) AS "count",
        category in (SELECT cat FROM checked_categories) as "selected"
    FROM items
    GROUP BY category
),
release_statuses_filtered AS (
    SELECT DISTINCT 'release_status' AS "name",
        release_status::text as "value",
        COUNT(release_status) AS "count"
    FROM items_filtered
    GROUP BY release_status 
),
release_statuses AS (
    SELECT DISTINCT 'release_status' AS "name",
        release_status::text as "value",
        COUNT(release_status) AS "count",
        release_status in (SELECT rs FROM checked_release_status) as "selected"
    FROM items
    GROUP BY release_status 
)
-- Union filters tables
SELECT a.name as "name", a.value, a.count, b.count as "count_filtered", a.selected FROM categories as a FULL JOIN categories_filtered as b ON a.name = b.name AND a.value = b.value
UNION
SELECT a.name as "name", a.value, a.count, b.count as "count_filtered", a.selected FROM release_statuses as a FULL JOIN release_statuses_filtered as b ON a.name = b.name AND a.value = b.value
ORDER BY "value" DESC
;