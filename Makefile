DC := docker compose -f docker-compose.yml

.PHONY: setup-db
setup-db:
	${DC} exec mysql bash -c 'mysql -uroot -ppassword < /app/setup.sql'

.PHONY: run-db
run-db:
	${DC} up -d mysql

.PHONY: run
run:
	${DC} up -d server

.PHONY: restart
restart:
	${DC} restart server

.PHONY: ps
ps:
	${DC} ps

