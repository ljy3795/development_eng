#!/bin/bash
docker run -d -p 8080:8080 --env-file .env --rm --name go-mls-bootcamp mls-bootcamp:server