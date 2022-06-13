DC := docker compose -f docker-compose.yml

.PHONY: setup-db
setup-db:
	${DC} exec mysql bash -c 'mysql -uroot -ppassword < /app/setup.sql'

.PHONY: ps
ps:
	${DC} ps

.PHONY: run
run:
	${DC} up -d server

.PHONY: restart
restart:
	${DC} restart server

