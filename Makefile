info-api:
	export PLANTINFO_DB_DSN='postgres://tester:aTestingPassword@localhost/plantinfo?sslmode=disable'
	docker-compose up -d database
	sleep 3
	go run ./cmd/api

info-web:
	docker-compose up -d database api
	go run ./cmd/web

everything:
	docker-compose up -d

clean:
	docker-compose down -v
	sudo rm -rf ./plantinfo-database/data

.PHONY: clean everything info-web info-api