package query

// PostgreSQL queries
// Drivers: postgres
const GET_TABLE_LIST_PSQL string = `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';`
const GET_TABLE_PK_PSQL string = `SELECT kcu.column_name FROM information_schema.table_constraints tc JOIN  information_schema.key_column_usage kcu  ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema WHERE tc.constraint_type = 'PRIMARY KEY' AND tc.table_name = '%s';`
const GET_TABLE_FKS_PSQL string = `SELECT tc.table_schema, tc.table_name, kcu.column_name, ccu.table_schema AS foreign_table_schema, ccu.table_name AS foreign_table_name, ccu.column_name AS foreign_column_name FROM information_schema.table_constraints AS tc JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema JOIN information_schema.constraint_column_usage AS ccu ON ccu.constraint_name = tc.constraint_name AND ccu.table_schema = tc.table_schema WHERE tc.constraint_type = 'FOREIGN KEY' AND tc.table_name = '%s';`
const GET_TABLE_RESTRAINS_PSQL string = `SELECT c.column_name, c.is_nullable, c.data_type, c.character_maximum_length, t.typname AS enum_type FROM information_schema.columns c JOIN pg_type t ON c.udt_name = t.typname WHERE c.table_name = '%s';`
const GET_TABLE_UNIQUE_COLS_PSQL string = `SELECT kcu.column_name FROM information_schema.table_constraints AS tc JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema WHERE tc.constraint_type = 'UNIQUE' AND kcu.table_name = '%s';`

const GET_ENUM_LIST_PSQL string = `SELECT t.typname AS enum_name, e.enumlabel AS enum_value FROM pg_type t JOIN pg_enum e ON t.oid = e.enumtypid JOIN pg_namespace n ON n.oid = t.typnamespace WHERE t.typcategory = 'E' AND n.nspname NOT IN ('pg_catalog', 'information_schema')  ORDER BY t.typname, e.enumsortorder;`

// MySQL queries
// Drivers: mysql, mariadb
const GET_TABLE_LIST_MYSQL string = `SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE();`
const GET_TABLE_PK_MYSQL string = `SELECT column_name FROM information_schema.key_column_usage WHERE table_schema = DATABASE() AND table_name = '%s' AND constraint_name = 'PRIMARY';`
const GET_TABLE_FKS_MYSQL string = `SELECT tc.constraint_schema AS table_schema, tc.table_name, kcu.column_name, kcu.referenced_table_schema AS foreign_table_schema, kcu.referenced_table_name AS foreign_table_name, kcu.referenced_column_name AS foreign_column_name FROM information_schema.table_constraints AS tc JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.constraint_schema = kcu.constraint_schema WHERE tc.constraint_type = 'FOREIGN KEY' AND tc.table_name = '%s';`
const GET_TABLE_RESTRAINS_MYSQL string = `SELECT c.column_name, c.is_nullable, c.data_type, CAST(c.character_maximum_length AS UNSIGNED) AS max_length, CASE WHEN c.data_type = 'enum' THEN SUBSTRING(c.column_type, 6, LENGTH(c.column_type) - 6 - 1) ELSE NULL END AS enum_type FROM information_schema.columns c WHERE c.table_name = '%s' AND c.table_schema = DATABASE();`
const GET_TABLE_UNIQUE_COLS_MYSQL string = `SELECT kcu.column_name FROM information_schema.table_constraints AS tc JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.constraint_schema = kcu.table_schema WHERE tc.constraint_type = 'UNIQUE' AND kcu.table_name = '%s' AND kcu.table_schema = DATABASE();`

// Sqlite queries
// Drivers: sqlite3
const GET_TABLE_LIST_SQLITE string = `SELECT name FROM sqlite_master WHERE type = 'table' AND name NOT LIKE 'sqlite_%';`
const GET_TABLE_PK_SQLITE string = `SELECT name FROM pragma_table_info('%s') WHERE pk > 0;`
const GET_TABLE_FKS_SQLITE string = `SELECT NULL as table_schema, NULL as table_name, "from" as column_name, NULL as foreign_table_schema, "table" as foreign_table_name, "to" as foreign_column_name FROM pragma_foreign_key_list('%s');`
const GET_TABLE_RESTRAINS_SQLITE string = `SELECT name AS column_name, CASE WHEN "notnull" = 0 THEN  'YES' ELSE 'NO' END as is_nullable, type AS data_type, NULL as character_max_length, NULL as enum_type FROM pragma_table_info('%s');`
const GET_TABLE_UNIQUE_COLS_SQLITE string = `WITH idx AS (SELECT name FROM pragma_index_list('%s') WHERE [unique] = 1) SELECT DISTINCT ii.name as column_name FROM idx il JOIN pragma_index_info(il.name) ii;`
