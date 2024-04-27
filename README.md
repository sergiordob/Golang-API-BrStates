# Golang-API-BrStates

This is a simple example of a REST API in Go (Golang) that uses version 1.22 of Go itself to create an HTTP server and the GORM ORM to interact with a PostgreSQL database.

## Configuration

### Dependencies
Make sure you have Go installed on your machine. You will also need to install the following dependency:

- [GORM](https://gorm.io/): Go ORM.

### Database
You can set up a PostgreSQL database using Docker Compose. The `docker-compose.yml` file is already configured to start two services: PostgreSQL and PgAdmin.

1. Make sure you have Docker and Docker Compose installed on your machine.
2. Run the following command in the root of the project: `docker-compose up`

This command will start the PostgreSQL and PgAdmin services.

Access PgAdmin in your browser using the following address: `http://localhost:54321`







