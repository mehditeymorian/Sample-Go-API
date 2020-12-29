#Internet Engineering Homework
Write API for a vendor with non-persistent and persistent database.



## How to Run
you can use Intellij IDE or [].exe files to run the project.

## Non-persistent Version
the default of project is non-persistent.
use ***API_NON_PERSISTENT.exe*** or Intellij IDE to run.

## Persistent Version
for the persistent version, you need a postgres server. after running your database, do the followings:

1. Configure **gorm** dsn such as host, user, pass and... information in **_db/PostgresDB.go_**.
2. Comment line 6 `var Database = CreateListDB()` in _**db/DBLayer.go**_
3. Uncomment line 5 `var Database = CreatePostgresDB()` in _**db/DBLayer.go**_
4. Run the server using one of the ways:
    1. Run the API using Intellij IDE.
    2. Execute `go build -o api.exe main.go` in project directory and run the api.exe

default configuration for postgres database is:
1. host = "localhost"
2. user = "postgres"
3. pass = "admin"
4. dbname = "customers"
5. port = "5432"

