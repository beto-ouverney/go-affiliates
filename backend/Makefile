

.SILENT:
.DEFAULT_GOAL := help

help:
	$(info falemais-telzir backend commands:)
	$(info -> test                    runs available tests)
	$(info -> run                     runs application)

test:
	echo "Running backend tests... To run all tests the docker-test must be running"
	go test -v ./...

run:
	echo "Running application... waiting a message:  Server running on port..."
	cd cmd && go run server.go

