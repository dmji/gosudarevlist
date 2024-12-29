-- name: InsertUpdateNoteItems :copyfrom
INSERT INTO animelayer_update_notes (update_id, title, value_old, value_new)
VALUES (@update_id, @title, @value_old, @value_new);