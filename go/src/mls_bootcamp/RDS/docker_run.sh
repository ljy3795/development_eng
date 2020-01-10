#!/bin/bash

docker run --name mariadbtest -e MYSQL_ROOT_PASSWORD='aa' -p 3306:3306 -d mariadb:latestdocker