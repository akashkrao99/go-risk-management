# Go Sample HTTP Server
This is a simple implementation of a simple http server using golang.

# How to run the project
Make sure you have `go 1.22` or higher installed on your system.

In order to execute the project locally using `make` run the following command:

```bash
make run_local
```

Alternatively if `make` does not work for you, please execute the following command:

```bash
go run cmd/main.go
```

To run the tests , please execute the following command:

```bash
cd internal/risks/__tests__
go test -v
```

To run the project using Docker Container, execute the following command:

```bash
docker compose up --build
```