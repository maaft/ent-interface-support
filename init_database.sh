#!/bin/bash
source server/.env && export $(cut -d= -f1 server/.env)

go run init/init-server.go