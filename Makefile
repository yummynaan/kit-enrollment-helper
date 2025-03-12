OS				:= $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH			:= $(shell uname -m | sed -e 's/x86_64/amd64/' -e 's/aarch64/arm64/')
MIGR_VER	:= v4.18.2
MIGRATE		:= ./migrate

init/tools: migrate/install

migrate/install:
	@if [ ! -x "$(MIGRATE)" ]; then\
		curl -L https://github.com/golang-migrate/migrate/releases/download/$(MIGR_VER)/migrate.$(OS)-$(ARCH).tar.gz | tar xvz --wildcards 'migrate';\
	fi
	$(MIGRATE) --version
	
		