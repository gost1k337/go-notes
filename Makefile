include .env
export

RUNNER=migrate

ifeq ($(p),host)
 	RUNNER=sql-migrate
endif

SOURCE="FILE://MIGRATIONS"
MIGRATE=$(RUNNER)
DB=${POSTGRES_DSN}
SWAGGER=swag
SWAGGER_PATH="./internal/app/app.go"

swagger init:
	$(SWAGGER) init -g ${SWAGGER_PATH}

migrate-status:
	$(MIGRATE) version

migrate-up:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" up 1

migrate-down:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" down 1

run dev:
	go run cmd/app/main.go
