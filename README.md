# PgBouncer Connection Demo

This project demonstrates using PgBouncer with Go's `database/sql` package to manage PostgreSQL connections efficiently.

## Features

- Connect to PostgreSQL via PgBouncer.
- Verify database connectivity.
- Execute a sample query.
- Simulate multiple concurrent connections to demonstrate PgBouncer's connection pooling.
- Gracefully close connections on application exit.

## Configuration

Set up your environment variables as follows:

	export DB_NAME="mydb"
	export DB_USER="myuser"
	export DB_PASSWORD="mypass"
	export DB_HOST="localhost"
	export PGBOUNCER_PORT="6432" # PgBouncer port


## Usage

Clone the repository:

    git clone https://github.com/your-username/pgbouncer-demo.git
    cd pgbouncer-demo

## Run the application:

    go run main.go

## Functions

    InitDatabase(): Initializes and verifies the PgBouncer connection.
    simulateConcurrentConnections(count int): Demonstrates pooling with concurrent queries.
    executeQuery(): Executes a sample query.
    Cleanup(): Ensures proper connection closure.

## Sample Output

Starting PgBouncer demonstration...
Database initialized successfully and PgBouncer connection verified.
Simulating multiple concurrent connections...
All concurrent connections completed.
PgBouncer demonstration completed.
Database connection closed.

## License

This project is licensed under the MIT License.
