plant-profile-api:
	docker-compose up -d database
	sleep 3
	export PLANTINFO_DB_DSN='postgres://tester:aTestingPassword@localhost/plantinfo?sslmode=disable' && \
	go run ./services/plant-profile-api/

tower-profile-api:
	docker-compose up -d database
	sleep 3
	export GROWTOWERINFO_DB_DSN='postgres://tester:aTestingPassword@localhost/growtowerinfo?sslmode=disable' && \
	go run ./services/tower-profile-api/

ui:
	docker-compose up -d database plantinfo-api towerinfo-api
	go run ./services/web

everything:
	docker-compose up -d

clean:
	docker-compose down -v
	sudo rm -rf ./plantinfo-database/data

.PHONY: clean everything ui plant-profile-api tower-profile-api