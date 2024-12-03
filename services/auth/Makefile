#ABOUT POSTGRES
DB_DRIVE = postgresql
DB_USER = postgres
DB_PASS = qwerty
DB_IP = localhost
DB_PORT = 5432
DB_NAME = postgres

#CONFIG
CONFIGPATH=./configs
CONFIGNAME=local
MIGRATIONSDIR=migrations

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
	migrate -path $(PATH_TO_SCHEMA) -database "$(DB_DRIVE)://$(DB_USER):$(DB_PASS)@$(DB_IP):$(DB_PORT)/$(DB_NAME)?$(DB_OPT)" -verbose up

migr_down: migr_up
	migrate -path $(PATH_TO_SCHEMA) -database "$(DB_DRIVE)://$(DB_USER):$(DB_PASS)@$(DB_IP):$(DB_PORT)/$(DB_NAME)?$(DB_OPT)" -verbose up

.DEFAULT_GOAL := run
.PHONY  : migr_down, migr_up, create_migration_files, run