PWD				:= $(shell pwd)
OS				:= $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH			:= $(shell uname -m | sed -e 's/x86_64/amd64/' -e 's/aarch64/arm64/')
MIGR_VER	:= v4.18.2
MIGRATE		:= $(PWD)/migrate
DB_PORT		:= 3308

init/tools: tbls/install migrate/install

tbls/install:
	@if [ ! -x "$(command -v tbls)" ]; then\
		go install github.com/k1LoW/tbls@latest;\
	fi
	tbls version

migrate/install:
	@if [ ! -x "$(MIGRATE)" ]; then\
		curl -L https://github.com/golang-migrate/migrate/releases/download/$(MIGR_VER)/migrate.$(OS)-$(ARCH).tar.gz | tar xvz --wildcards 'migrate';\
	fi
	$(MIGRATE) --version
	
migrate/create:
	$(MIGRATE) create -ext sql -dir migrations -seq $(FILENAME)

compose/database/init: compose/database/up sleep docker/mysql/migrate

compose/database/up:
	docker compose up -d

compose/database/down:
	docker compose down 

sleep:
	until (mysqladmin ping -h 127.0.0.1 -P $(DB_PORT) -uroot -proot --silent) do echo 'waiting for mysql connection...' && sleep 2; done

docker/mysql/migrate:
	$(MIGRATE) --path migrations --database 'mysql://root:root@tcp(localhost:$(DB_PORT))/kit_enrollment_helper?parseTime=true' up
