-- name: GetFilters :many
WITH checked_release_status AS (
    SELECT unnest(
            ARRAY [@checked_status_array::RELEASE_STATUS_ANIMELAYER[]] 
    ) as rs
),
selected_release_status as (
     SELECT unnest(
            ARRAY [@selected_status_array::RELEASE_STATUS_ANIMELAYER[]] 
    ) as rs
),
-- Collect filtered items
items_filtered AS (
    SELECT release_status,
        created_date,
        updated_date
    FROM animelayer_items
    WHERE release_status IN (SELECT rs FROM selected_release_status)
        AND (
            @search_query::text = ''
            OR SIMILARITY(title, @search_query) > @similarity_threshold::float
        )
),
-- Collect filters table
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
    FROM animelayer_items
    GROUP BY release_status 
)
-- Union filters tables
SELECT a.name::text as "name", a.value::text as "value", a.count::bigint as "count", COALESCE(b.count, 0)::bigint as "count_filtered", a.selected::boolean as "selected" FROM release_statuses as a FULL JOIN release_statuses_filtered as b ON a.name = b.name AND a.value = b.value
UNION
SELECT a.name::text as "name", a.value::text as "value", a.count::bigint as "count", COALESCE(b.count, 0)::bigint as "count_filtered", a.selected::boolean as "selected" FROM release_statuses as a FULL JOIN release_statuses_filtered as b ON a.name = b.name AND a.value = b.value
ORDER BY "value" DESC;