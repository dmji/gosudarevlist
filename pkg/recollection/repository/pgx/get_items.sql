-- name: GetItems :many
SELECT *
FROM animelayer_items
WHERE category = @category::CATEGORY_ANIMELAYER
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
LIMIT @count OFFSET @offset_count;