DC := docker compose -f docker-compose.yml

.PHONY: ps
ps:
	${DC} ps


.PHONY: run
run:
	${DC} up -d server

.PHONY: restart
restart:
	${DC} restart server

