# go-clean-arch-sample-app

TheApp is a sample application implemented in Golang using Domain-Driven Design (DDD) and Clean Architecture principles.

## Layered Architecture Overview

```
├── constants
│   └── constants.go
├── domain
│   ├── model
│   ├── repository
│   └── service
├── infrastructure
│   └── datastore
├── interfaces
│   └── commandline
└── usecase
```


Each layer has a specific responsibility:

- **Domain**: Contains the core business logic, independent of external systems.
- **Usecase**: Implements application-specific business rules, orchestrating interactions between the domain and other layers.
- **Interfaces**: Handles input/output, adapting external systems (e.g., CLI, web) to the application.
- **Infrastructure**: Manages external dependencies like databases or third-party services.

This structure ensures a clear separation of concerns and makes the application easier to test, maintain, and extend.

## Requirements

- Go 1.23.1 or later

## Usage

To run the application, use the following command:

```sh
go run main.go
```