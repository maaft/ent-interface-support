#!/bin/bash

source server/.env && export $(cut -d= -f1 server/.env)

(cd server/.docker; docker-compose up &)

go run server/server.go