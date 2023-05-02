-- these columns need indexes because they're in fk
select tc.constraint_type, kcu.* from information_schema.table_constraints tc inner join information_schema.key_column_usage kcu
on tc.constraint_name = kcu.constraint_name
where tc.table_schema = 'world' and tc.constraint_type = 'FOREIGN KEY';