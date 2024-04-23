# Simple API using Go
Here is a simple RESTful API using Go, Serverless, AWS API Gateway, Lambda, and DynamoDB.

# Quick start
To set up project follow these steps:

- First install all module dependencies
```bash 
go mod tidy
```

- Up and run Database. If you execute project locally, setup DynamoDB via docker. Otherwise, on AWS you need to define Instance.
```
docker compose up
```

- Execute bash script to define DynamoDB table and seed data at local.
```
./schema/schema-seed-data.sh
```

- To start the server use the following command:
```
go run main.go
```
Finally,  open http://localhost:8080

# Folder structure
This project use common folder structure for a Go REST API:

```bash

simple-api/
├── main.go
├── handlers/
│   └── device_handler.go
├── routes/
│   └── routes.go
├── models/
│   └── device.go
├── repositories/
│   ├── device_repository.go
│   └── device_memory_repository.go
│   └── device_dynamodb_repository.go
├── services/
│   └── device_service.go
├── db/
│   └── db.go
└── utils/
    └── utils.go
```
Here's a breakdown of the folders and files:

- `main.go`: This is the entry point of application, where is initialized and run server, as well as set up dependencies and configurations.
- `handlers/device_handler.go`: This file contains the HTTP handler functions for the Device resource, handling the CRUD operations.
- `routes/routes.go`: This file defines the routes for API, including the Device resource routes.
- `models/device.go`: This file defines the `Device` struct and any related types or methods.
- `repositories/device_repository.go`: This is an interface that defines the methods for interacting with the Device data store.
- `repositories/device_memory_repository.go`: This is an in-memory implementation of the `DeviceRepository` interface.
- `repositories/device_dynamodb_repository.go`: This is an DynamoDB implementation of the `DeviceRepository` interface.
- `services/device_service.go`: This file contains the business logic for the Device resource, orchestrating the interactions between the repository and the handlers.
- `db/db.go`: This file sets up the database connection and provides a way to access the database object throughout the application.
- `utils/utils.go`: This folder can hold any reusable utility functions or packages used across the application.

This structure is flat and straightforward, with fewer nested folders. It's a common and widely accepted structure for smaller to medium-sized projects.

# Testing
From the project’s root directory, run test:
```bash
go test -v ./...

```

To view the test coverage you can run:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out 
```

# TODO
- [ ] Document Project specifications
- [ ] Document Configuration
- [ ] Deploy on AWS

 
