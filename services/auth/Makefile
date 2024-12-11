#ABOUT POSTGRES
DB_DRIVE = postgres
DB_USER = postgres
DB_PASS = qwerty
DB_IP = localhost
DB_PORT = 5432
DB_NAME = postgres
DB_OPT = sslmode=disable

#CONFIG
CONFIGPATH = ./configs
CONFIGNAME = local

#DATABASE
MIGRATIONSDIR = migrations
TEST_MIGRATIONSDIR = tests/migration_test

build :
	go build -o ./.bin/auth ./src/cmd/main.go

run : build
	./.bin/auth  --config_path=$(CONFIGPATH) --config_name=$(CONFIGNAME)
# 	 go run src/cmd/main.go --config=$(CONFIGPATH)
clean :
	rm -rf ./.bin

create_migration_files :
	mkdir -p $(MIGRATIONSDIR)
	migrate create -ext sql -dir $(MIGRATIONSDIR) -seq init

# up_database_in_container :
# 	docker run 

migr_up:
	migrate -path $(MIGRATIONSDIR) -database "$(DB_DRIVE)://$(DB_USER):$(DB_PASS)@$(DB_IP):$(DB_PORT)/$(DB_NAME)?$(DB_OPT)" -verbose up

migr_down: migr_up
	migrate -path $(MIGRATIONSDIR) -database "$(DB_DRIVE)://$(DB_USER):$(DB_PASS)@$(DB_IP):$(DB_PORT)/$(DB_NAME)?$(DB_OPT)" -verbose up

test_create_migration_files:
	migrate create -ext sql -dir $(TEST_MIGRATIONSDIR) -seq init

test_migr_up:
	migrate -path $(TEST_MIGRATIONSDIR) -database "$(DB_DRIVE)://$(DB_USER):$(DB_PASS)@$(DB_IP):$(DB_PORT)/$(DB_NAME)?$(DB_OPT)" -verbose up

test_migr_down: migr_up
	migrate -path $(TEST_MIGRATIONSDIR) -database "$(DB_DRIVE)://$(DB_USER):$(DB_PASS)@$(DB_IP):$(DB_PORT)/$(DB_NAME)?$(DB_OPT)" -verbose down 

drop_migr : 
	docker exec -it postgres_auth /bin/bash 
# psql -U postgres
# UPDATE schema_migrations SET dirty=false WHERE version=1;
# DROP TABLE IF EXISTS schema_migrations;

.DEFAULT_GOAL := run
.PHONY  : migr_down, migr_up, create_migration_files, run