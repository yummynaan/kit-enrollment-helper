OS				:= $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH			:= $(shell uname -m | sed -e 's/x86_64/amd64/' -e 's/aarch64/arm64/')
MIGR_VER	:= v4.18.2
MIGRATE		:= ./migrate

init/tools: sql-migrate/install

sql-migrate/install:
	@if [ ! -x "$(command -v sql-migrate)" ]; then\
		go install github.com/rubenv/sql-migrate/...@latest;\
	fi
	sql-migrate --version

golang-migrate/install:
	@if [ ! -x "$(command -v migrate)" ]; then\
		curl -L https://github.com/golang-migrate/migrate/releases/download/$(MIGR_VER)/migrate.$(OS)-$(ARCH).tar.gz | tar xvz --wildcards 'migrate';\
	fi
	$(MIGRATE) --version
	
		