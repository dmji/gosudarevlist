/* -- name: InsertNewUpdateNote :exec
 INSERT INTO animelayer_updates (
 item,
 update_date,
 title,
 value_old,
 value_new
 )
 VALUES (
 @item,
 @update_date,
 @title,
 @value_old,
 @value_new
 ); */