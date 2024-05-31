# SOMOS Events

The backend for [SOMOS](https://somos831.com/).

## Getting started:

Install golang:
https://go.dev/doc/install

Install mysql: 
https://dev.mysql.com/downloads/mysql/

You can install mysql workbench (if needed) from:
https://dev.mysql.com/downloads/

Also install the golang-migrate cli tool listed bellow.

### DB Migrations:
We are using golang-migrate for our database migrations.

- Install the CLI tool: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

- Confirm installation: `migrate -version`

More information about the way we create migrations can be found [here](https://github.com/somos831/somos-backend/tree/main/db).

### Running the project locally: 

In the root of the project there is a file called `.example.env`. This is an example configuration file, make a copy of it and name it `.env`. 
Then update the environment variable values to match your local environment's creadentials.

Once that is done you can run `make run` in the projects root directory to start running the application. This will run any up migrations and start the server.