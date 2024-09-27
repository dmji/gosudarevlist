-- name: InsertNewDescriptionNote :exec
INSERT INTO animelayer_description_notes (description_id, field_name, field_text)
VALUES (
        @description_id,
        @field_name,
        @field_text
    );

-- name: UpdateDescriptionNote :exec
UPDATE animelayer_description_notes
SET field_text = @field_text
WHERE description_id = @description_id
    AND field_name = @field_name;