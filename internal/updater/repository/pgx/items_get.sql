-- name: GetItems :many
SELECT *
FROM animelayer_items
WHERE last_checked_date > @last_checked_date;