create table user_audit
(
    id            int generated always as identity primary key not null,
    op        varchar(10)                                  not null,
    changed_at    timestamp,
    user_name          varchar(50)                                  not null,
    email         varchar(50)                                  not null
);

create function user_note_changes() returns trigger as $$
begin
    IF TG_OP='INSERT' then
        insert into user_audit(op, changed_at, user_name, email)
        select 'insert', now(), name, email from NEW;
ELSEIF TG_OP='UPDATE' then
        insert into user_audit(op, changed_at, user_name, email)
        select 'update', now(), name, email from NEW;
ELSEIF TG_OP='DELETE' then
        insert into user_audit(op, changed_at, user_name, email)
        select 'delete', now(), name, email from OLD;
end if;
return null;
end $$ language plpgsql;


drop trigger if exists audit_user_insert on users;
create trigger audit_user_insert after insert on users
referencing new table as new
for each statement execute procedure user_note_changes();

drop trigger if exists audit_user_update on users;
create trigger audit_user_update after update on users
referencing new table as new
for each statement execute procedure user_note_changes();

drop trigger if exists audit_user_delete on users;
create trigger audit_user_delete after delete on users
referencing old table as old
for each statement execute procedure user_note_changes();