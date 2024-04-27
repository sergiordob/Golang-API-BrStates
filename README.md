# Golang-API-BrStates

This is a simple example of a REST API in Go (Golang) that uses version 1.22 of Go itself to create an HTTP server and the GORM ORM to interact with a PostgreSQL database.

'Acre'
'Alagoas'
'Amapá'
'Amazonas'
'Bahia'
'Ceará'
'Distrito Federal'
'Espírito Santo'
'Goiás',
'Maranhão', 
'Mato Grosso',
'Mato Grosso do Sul',
'Minas Gerais',
'Pará', 
'Paraíba',
'Paraná', 
'Pernambuco', 
'Piauí',
'Rio de Janeiro',
'Rio Grande do Norte'
'Rio Grande do Sul', 
'Rondônia', 
'Roraima',
'Santa Catarina', 
'São Paulo', 
'Sergipe',
'Tocantins'

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

Use the email eletricamat@gmail.com and password admin to log in. Then, you can add a PostgreSQL server and connect to the database.

## Migration Script

During the initialization of the PostgreSQL service, Docker Compose will run an SQL script located at `./migration/docker-database-initial.sql.` This script is responsible for creating the necessary database and table for the application.

## How to Use

1. Clone this repository to your local machine.
2. Make sure you have the GORM dependency installed.
3. Set up the Docker Compose environment as described above.
4. Run the server with the following command: `go run main.go`

## Endpoints

1. `GET /api`: Returns all states in JSON format.
2. `GET /api/{id}`: Returns a specific state by ID in JSON format.  
3. `GET /api/total`: Returns all states in JSON format.

## Project Structure

1. `config/`: Contains the HTTP server configuration.
2. `controllers/`: Controllers to handle HTTP requests.
3. `database/`: Configuration and connection with the database.
4. `models/`: Definition of the application data models.
5. `routes/`: Definition of API routes.
6. `main.go`: Entry point of the application, where routes are loaded and the server is initialized.

