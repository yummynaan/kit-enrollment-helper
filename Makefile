init/tools: sql-migrate/install

sql-migrate/install:
	@if [ ! -x "$(command -v sql-migrate)" ]; then\
		go install github.com/rubenv/sql-migrate/...@latest;\
	fi
	sql-migrate --version