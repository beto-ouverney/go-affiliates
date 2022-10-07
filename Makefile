# Definitions
BACKEND := backend
FRONTEND := frontend
CONTAINER_BACK_TEST := affiliates_backend_test
CONTAINER_FRONT_TEST := affiliates_frontend_test
CONTAINER_BACK := affiliates_backend
CONTAINER_FRONT := affiliates_frontend
DOCKER_COMPOSE_TEST   := docker-compose.test.yml
DOCKER_COMPOSE_DEV    := docker-compose.dev.yml

.SILENT:
.DEFAULT_GOAL := help

help:
	$(info falemais-affiliates commands:)
	$(info -> docker-test-up          up docker for test)
	$(info -> docker-test-down        down docker for test)
	$(info -> docker-up               up docker for production)
	$(info -> docker-down             down docker for production)
	$(info -> back-test               runs backend tests)
	$(info -> run-back                runs backend application)
	$(info -> run-front               runs frontend application)
	$(info -> back-container-test     access backend container test bash)
	$(info -> back-container          access backend container bash)
	$(info -> front-container-test    access front container bash)
	$(info -> npm                     to install npm dependencies)

docker-test-up:
	docker-compose -f ${DOCKER_COMPOSE_TEST} up -d --build

docker-test-down:
	docker-compose -f ${DOCKER_COMPOSE_TEST} down

docker-up:
	docker-compose -f ${DOCKER_COMPOSE_DEV} up -d --build

docker-down:
	docker-compose -f ${DOCKER_COMPOSE_DEV} down

back-container-test:
	docker exec -it ${CONTAINER_BACK_TEST} bash

back-container:
	docker exec -it ${BACKEND} bash

front-container-test:
	docker exec -it ${CONTAINER_FRONT_TEST} sh

back-test:
	echo "Running backend tests... Remender, to run all tests the docker-test must be running"
	cd ${BACKEND} && go test -v ./...

run-back:
	echo "Running application... waiting a message:  Server running on port..."
	cd ${BACKEND} && go run server.go

npm:
	echo "Install dependences frontend application..."
	cd ${FRONTEND} && npm install

run-front:
	echo "Running application frontend application..."
	cd ${FRONTEND} && npm run serve
