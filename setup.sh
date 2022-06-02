#!/bin/bash
docker-compose exec mysql bash -c 'mysql -uroot -ppassword < /app/setup.sql'
