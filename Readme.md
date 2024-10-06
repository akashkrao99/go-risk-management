# Go Sample HTTP Server
This is a simple implementation of a simple http server using golang.

# How to run the project
Make sure you have `go 1.22` or higher installed on your system.

To ensure all dependencies are downloaded, run the following command:

```bash
go mod download
```

In order to execute the project locally using `make` run the following command:

```bash
make run_local
```

Alternatively if `make` does not work for you, please execute the following command:

```bash
go run cmd/main.go
```

Alternatively, to run the project using Docker Container, execute the following command:

```bash
docker compose up --build
```

## Test the APIs locally

After your local environment is running, you can import the demo Postman collection from [here](./RisksDemo.postman_collection.json) into Postman and use it to test the APIs.

## Executing Testcases

To run the tests , please execute the following command:

```bash
cd internal/risks/__tests__
go test -v
```