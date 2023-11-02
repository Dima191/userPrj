drop trigger if exists audit_user_insert on users;
drop trigger if exists audit_user_update on users;
drop trigger if exists audit_user_delete on users;
drop function if exists user_note_changes;
drop table if exists user_audit;