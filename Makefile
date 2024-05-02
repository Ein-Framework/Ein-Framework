.PHONY: build
build:
	go build -o ein ./cmd/ein.go

.PHONY: install
install:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

.PHONY: run
run:
	air -c server.air.toml

.PHONY: clean
clean:
	rm -f main

.PHONY: docker-build
docker-build:
	docker-compose -f ./scripts/docker-compose.yml build 

.PHONY: docker-run
docker-run:
	docker-compose -f ./scripts/docker-compose.yml up
