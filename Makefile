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

.PHONY: test
test:
	go test -race -vet=off -v ./...

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
	docker-compose -f ./scripts/db.docker-compose.yml up -d

.PHONY: test-cov
test-cov:
	go test -coverprofile ./test.cov -race -vet=off -v ./...
	go tool cover -html=test.cov -o test.cov.html

.PHONY: build-plugins
build-plugins:
	go build -o ./plugins/dist/http.so -buildmode=plugin ./plugins/http/main.go

.PHONY: cp-plugins
cp-plugins:
	cp ./plugins/dist/* ~/.ein-framework/plugins/

.PHONY: install-plugins
install-plugins: build-plugins cp-plugins

.PHONY: install-templates
install-templates:
	cp ./templates/* ~/.ein-framework/templates/

.PHONY: setup-framework
setup-framework: install-plugins install-templates
