# ein-framework

- Go version: Go 1.22.1

- Build binary

```shell
go build -o ein .
```

- Start Development workflow

```shell
	air
```

- Clean workspace

```shell
rm -f ein
```

- Build Docker

```
docker-compose -f ./scripts/docker-compose.yml build
```

- Run Docker

```shell
	docker-compose -f ./scripts/docker-compose.yml up
```
