## Database Migrations
We are using [golang-migrate](https://github.com/golang-migrate/migrate) to manage database schema changes efficiently and reliably in our Go project. **golang-migrate** provides a straightforward and version-controlled approach to database migrations, ensuring that database changes are seamlessly applied across different environments and enabling us to maintain a consistent and up-to-date database schema throughout the development lifecycle.

### Getting started with golang-migrate:
- Install the CLI tool, instructions can be found here: 
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

- Verify it has been installed: `migrate -version`

- For local development, ensure that you have an instance of MySQL and that you have created a database for this project. Follow the .example.env file to create your own `.env` file with the corresponding credentials to your database.

- Once you have your database running and your `.env` file populated you can run `make` in the root of the project. That will run any migrations and run the application.

### Creating new migrations:
We will be using sequential numbers to version our database migrations. Adding migrations is simple, they can be created by running the following command in the root of the project:
```
migrate create -ext sql -dir db/migrations -seq description-of-schema-change
```

This will create two files under `db/migrations`. One is the up migration and the other the equivalent down migration. Fill those files out and you'll be able to run those migrations.

To run **up** migrations:

```migrate -database "mysql://username:password@tcp(host:port)/database" -path "path/to/migrations" up```

To run **down** migrations:

```migrate -database "mysql://username:password@tcp(host:port)/database" -path "path/to/migrations" down```