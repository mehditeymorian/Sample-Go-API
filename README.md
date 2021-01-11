# Go API
Write API for a vendor with non-persistent and persistent database. 
this was one of the internet engineering lecture homeworks. [Link to Homework Description](https://github.com/mehditeymorian/Sample-Go-API/blob/master/hw4.pdf)



## How to Run
you can use Intellij IDE or [].exe files to run the project. the default of project is non-persistent.

## Non-persistent Version
do the followings:
1. Uncomment line 6 `var Database = CreateListDB()` in _**db/DBLayer.go**_
2. Comment line 5 `var Database = CreatePostgresDB()` in _**db/DBLayer.go**_
3. Run the API using one of the ways:
    1. Using Intellij IDE.
    2. Execute `go build -o api.exe main.go` in project directory and run the api.exe

## Persistent Version
for the persistent version, you need a postgres server. after running your database, do the followings:

1. Configure **gorm** dsn such as host, user, pass and... information in **_db/PostgresDB.go_**.
2. Comment line 6 `var Database = CreateListDB()` in _**db/DBLayer.go**_
3. Uncomment line 5 `var Database = CreatePostgresDB()` in _**db/DBLayer.go**_
4. Run the API using one of the ways:
    1. Using Intellij IDE.
    2. Execute `go build -o api.exe main.go` in project directory and run the api.exe

default configuration for postgres database is:
1. host = "localhost"
2. user = "postgres"
3. pass = "admin"
4. dbname = "customers"
5. port = "5432"

