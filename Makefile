generate: generate_schema generate_gql_server

generate_schema:
	go generate ./ent/...

generate_gql_server:
	go generate ./...

run:
	bash run.sh

init_database: 
	bash init_database.sh

init_testdata: init_database
	bash test/init_testdata.sh

reset_db:
	(cd server/.docker/ && docker-compose down -v)