# EniQilo Store

Welcome to the EniQilo Store Service. An application where . This service has implemented clean architecture principles, a practical software architecture solution from Robert C. Martin (known as Uncle Bob).

## What's New?

The following are some of the updates in this latest version of the content management system:

- Standardize the response format.
- Refactor algorithm including handler, usecase, and repository layer.
- Add custom middlewares and validator.
- Update database schema.
- Minor bug fixes.

## Getting Started

### Prerequisites

- [Go 1.19.3](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

### Installation

- Clone the git repository:

```
$ git clone https://github.com/Assyatier21/eniqilo-store.git
$ cd eniqilo-store
```

- Install Dependencies

```
$ go mod tidy
```

- Adjust Configuration

```
$ cp .env.example .env
```

Then update config file as needed. I think you are quite easy for a smart person like you:)

### Running

```
$ go run cmd/main.go
```

or simply

```
$ make run
```

## API Reference

### Unit Testing

```
$ go test -v -coverprofile coverage.out ./...
```

## Install Local Sonarqube

please follow this [tutorial](https://techblost.com/how-to-setup-sonarqube-locally-on-mac/) as well.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/backend-magang/eniqilo-store/blob/master/LICENSE) file for details.
