.PHONY: build
build:
	go build -o ein ./cmd/ein.go

.PHONY: install
install:
	echo "Please run the install script with sudo privileges"
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
	curl https://raw.githubusercontent.com/llorllale/go-gitlint/master/download-gitlint.sh > download-gitlint.sh && bash download-gitlint.sh 
	mv ./bin/gitlint /usr/bin
	rm -rf ./bin/gitlint download-gitlint.sh
	echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
	sudo apt update
	sudo apt install goreleaser
	go install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: dev
dev:
	air 

.PHONY: release
release:
	goreleaser release --snapshot --clean

.PHONY: clean
clean:
	rm -f main
	rm -rf ./dist

.PHONY: docker-build
docker-build:
	docker-compose -f ./scripts/docker-compose.yml build 

.PHONY: docker-run
docker-run:
	docker-compose -f ./scripts/docker-compose.yml up

.PHONY: start-db
start-db:
	docker-compose -f ./scripts/db.docker-compose.yml up
