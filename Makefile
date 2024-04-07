.PHONY: build
build:
	go build -o ein .

.PHONY: run
run:
	air

.PHONY: clean
clean:
	rm -f main

.PHONY: docker-build

docker-build:
	docker-compose -f ./scripts/docker-compose.yml build 

.PHONY: docker-run
docker-run:
	docker-compose -f ./scripts/docker-compose.yml up
