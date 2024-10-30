DOCKER_COMPOSE_FILE ?= docker-compose.yml #「?=」の意味は変数が未定義のときのみ代入
POSTGRES_USER ?= postgres
POSTGRES_PASSWORD ?= secret
POSTGRES_DB=hands_on

all:
	docker compose up

up-web:
	docker compose up web

up-back:
	docker compose up backend

up-db:
	docker compose up postgres

down:
	docker-compose down

mig-up:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up

mig-down:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down

arg-check:
ifndef name
	@echo "エラー: 'name'が設定されていません。使用法: make create-create-sql name=テーブル名"
	exit 1
endif

create-create-sql: arg-check
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migration create_table_$(name)

create-alter-sql: arg-check
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migration alter_table_$(name)

create-trigger-sql: arg-check
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migration trigger_table_$(name)

add-init-sql: arg-check
	@SEED_FILE_PATH=./db/seed/$(name).csv; \
	if [ ! -f $$SEED_FILE_PATH ]; then \
 		cp ./db/migration/99999999999999_init_data_insert.up.sql ./db/migration/99999999999999_init_data_insert.up.sql.tmp; \
 		echo "COPY $(name) FROM '/seed/$(name).csv' DELIMITER ',' CSV HEADER;" >> ./db/migration/99999999999999_init_data_insert.up.sql.tmp; \
 		mv ./db/migration/99999999999999_init_data_insert.up.sql.tmp ./db/migration/99999999999999_init_data_insert.up.sql; \
 		touch $$SEED_FILE_PATH; \
 		echo "シードファイルを作成しました: $$SEED_FILE_PATH"; \
	else \
 		echo "seedファイルが既に存在します: $$SEED_FILE_PATH"; \
	fi

shell-db:
	docker compose -f ${DOCKER_COMPOSE_FILE} exec postgres psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

column_name: arg-check
	docker compose -f ${DOCKER_COMPOSE_FILE} exec postgres psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} \
	-c "SELECT string_agg(column_name, ',') AS columns FROM information_schema.columns WHERE table_name = '$(name)';"

