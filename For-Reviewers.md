# Go HTTP Server - Risk Management API

This is a sample HTTP server implemented in Go that provides a simple Risk Management API. It is designed with a layered architecture to separate concerns and demonstrate a clean, maintainable structure. The API exposes four endpoints for managing risks, including a health check endpoint.

## Table of Contents

- [Overview](#overview)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Thought Process](#thought-process)

## Overview

The project implements a basic risk management service that runs on port `8080`. It follows a clean architecture with the following layers:

- **Controller Layer**: Handles the HTTP requests and responses.
- **Service Layer**: Contains the business logic for managing risks.
- **Repository Layer**: Manages the interaction with the data source (which is an in-memory store in this case).

### API Endpoints

#### 1. **GET /v1/risks**

Fetches all risks available in the system.

- **Method**: `GET`
- **Endpoint**: `/v1/risks`
- **Description**: Retrieves a list of all risks stored in the system.
- **Response**:
  - `200 OK`: Returns the list of risks.
  - `500 Internal Server Error`: If any server-side issue occurs.

#### 2. **POST /v1/risks**

Creates a new risk entry.

- **Method**: `POST`
- **Endpoint**: `/v1/risks`
- **Description**: Adds a new risk to the system.
- **Request Body**:
  - `application/json`: Contains the details of the risk to be created.
  - Request Body
   ```json
   {
    "title": "Title of the risk", //mandatory
    "description": "Description of the risk", //mandatory
    "status": "status of the risk" //mandatory(can be among [open, closed, accepted, investigating])
  }
  ```
- **Response**:
  - `201 Created`: Successfully created a new risk.
  - `400 Bad Request`: If the input is invalid.
  - `500 Internal Server Error`: If any server-side issue occurs.

#### 3. **GET /v1/risks/{id}**

Fetches a specific risk by its ID.

- **Method**: `GET`
- **Endpoint**: `/v1/risks/{id}`
- **Description**: Retrieves details of a specific risk identified by its unique `id`.
- **Response**:
  - `200 OK`: Returns the risk details.
  - `404 Not Found`: If the risk does not exist.
  - `500 Internal Server Error`: If any server-side issue occurs.

#### 4. **GET /health**

Checks the health status of the service.

- **Method**: `GET`
- **Endpoint**: `/health`
- **Description**: Returns the health status of the API service.
- **Response**:
  - `200 OK`: Indicates the service is running properly.
  - `500 Internal Server Error`: If any health issue is detected.

## Project Structure

The project is divided into multiple layers, which provide separation of concerns and modularity:

1. **Controller Layer**: This layer is responsible for processing the incoming HTTP requests, invoking the appropriate service methods, and formatting the responses.

2. **Service Layer**: This layer encapsulates the business logic. It is responsible for applying the business rules and orchestrating data flow between the repository and controller.It is also responsible for handling basic validation and sanitization before passing the data on to the repository layer.

3. **Repository Layer**: This layer interacts with the underlying data source. In this example, it uses an in-memory store to save, retrieve, and update risks.

## Thought Process

The goal was to design a simple yet extensible HTTP server that provides basic CRUD-like functionality for managing risks. The process followed:

##### 1. Setup a Simple Router :

   - We started by setting up a Go HTTP server using the `gin` package and routing the endpoints using a simple router. The server runs on port `8080` by default.

##### 2. Health Check:

   - A `/health` endpoint was implemented to monitor the service status. This helps to confirm that the server is running without issues.

##### 3. Model for Request and Response:

   - We added models for handling risk requests and responses. This helps enforce basic validation and sanitization at the model level, ensuring the data is clean and well-structured. It also provides consistency in the response format across different API endpoints.

##### 4. Middleware:

   - A middleware was introduced to add flexibility in handling common tasks. Currently, it is used to avoid CORS errors and blacklist certain IPs.

##### 5. In-Memory Database:

   - The repository layer uses a map as the in-memory data store. This allows for O(1) lookup times and simplicity in managing data during development. While sufficient for this example, it can be replaced with a persistent data store in the future .Considering the scope of this example we can also use redis.

##### 6. Configuration Management:
   - We initialized a configuration management system that reads variables from an environment file. A singleton pattern was implemented here to ensure that the configuration is not initialized multiple times throughout the application, providing a single source of truth for configuration values.
