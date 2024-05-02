# ein-framework

<<<<<<< HEAD
- Go version: Go 1.22.1
=======
> Go version: Go 1.22.1
>>>>>>> fd1645c99c3bd5bf3fb4fc33e66e20880321ed56

- Install env

```shell
<<<<<<< HEAD
make install
=======
sudo make install
>>>>>>> fd1645c99c3bd5bf3fb4fc33e66e20880321ed56
```

- Build binary

```shell
<<<<<<< HEAD
go build -o ein .
=======
make build
>>>>>>> fd1645c99c3bd5bf3fb4fc33e66e20880321ed56
```

- Start Development workflow

```shell
<<<<<<< HEAD
	air
=======
make dev 
>>>>>>> fd1645c99c3bd5bf3fb4fc33e66e20880321ed56
```

- Clean workspace

```shell
<<<<<<< HEAD
rm -f ein
=======
make clean
>>>>>>> fd1645c99c3bd5bf3fb4fc33e66e20880321ed56
```

- Build Docker

```
<<<<<<< HEAD
docker-compose -f ./scripts/docker-compose.yml build
=======
make docker-build
>>>>>>> fd1645c99c3bd5bf3fb4fc33e66e20880321ed56
```

- Run Docker

```shell
<<<<<<< HEAD
	docker-compose -f ./scripts/docker-compose.yml up
```
=======
make docker-run
```


- Build for release

```shell
make release
```
>>>>>>> fd1645c99c3bd5bf3fb4fc33e66e20880321ed56
