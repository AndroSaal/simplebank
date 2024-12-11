UPDATE schema_migrations SET dirty=false WHERE version=1;
-- DROP TABLE IF EXISTS schema_migrations;
DROP TABLE IF EXISTS migration_tests;