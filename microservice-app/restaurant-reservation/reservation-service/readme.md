## Description

Farmish CRM - REST application in Go.

## Installation

You cloned the repository! Several steps left to launch the server. All commands were written in Makefile

1. Create database and change DB_URL variable to your local postgres URL in Makefile. for example:

```bash
DB_URL=postgresql://postgres:root@localhost:5432/farmish?sslmode=disable
```

2. Create .env file in the root directory, copy all the vars from this example and fill them with required values.

```bash
PG_PASSWORD=root
SIGNING_KEY=farmish1234
```

3. Once you configured your database, run migration command to create tables:

```bash
$ make migrate-up
```

4. Install packages:

```bash
$ make tidy
```

### Running the app

```bash
# starts the server
$ make run
```

## Migration

Once you get into production you'll need to synchronize model changes into the database. Here is where migrations come to help:

```bash
# Create a new migration with specific name
$ make migrate-create NAME="YOUR_MIGRATION_NAME"
```

```bash
# Apply migration
$ make migrate-up
```

```bash
# Reset migrations
$ make migrate-down
```

## Documentation

To get acquainted with the list of existing endpoints and find out how it works go to _http://localhost:3000/swagger/index.html#/_ where you can find Swagger documentation.

To generate Swagger files:

```bash
$ make swag
```
