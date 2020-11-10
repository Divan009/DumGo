# Dummy Go App using GQLGEN

Go is used to create scalable infra. Learn more about Go here[https://golang.org/doc/install]

With these lines I have the most basic structure of Go Server running
1.	go mod init github.com/Divan009/gqlgen-todos
2.	go get github.com/99designs/gqlgen
3.	go run github.com/99designs/gqlgen init
4.	go run ./server.go

Whenever you chance the schema file, use:
go run github.com/99designs/gqlgen generate

To understand more about gqlgen.yml: https://gqlgen.com/config/
Facing issue with migrate(github.com/golang-migrate/migrate v3.5.4+incompatible)
Use go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/

JWT go get github.com/dgrijalva/jwt-go


#### Create a postgres db using pdAdmin or psql, whichever you prefer
CREATE TABLE IF NOT EXISTS Users(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Username VARCHAR (127) NOT NULL UNIQUE,
    Password VARCHAR (127) NOT NULL,
    PRIMARY KEY (ID)
)

CREATE TABLE IF NOT EXISTS Links(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Title VARCHAR (255) ,
    Address VARCHAR (255) ,
    UserID INT ,
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,
    PRIMARY KEY (ID)
)

## All the imp files explained

Schema.graphqls - basically where i define the structure or what is to be returned from API. This is the first thing I do.

Model_gen - This gets generated from the above schema.

Schema_resolver - for all your functions

Postgres - I create Db and call it

pqueries - is where i define db functions

getLogic - uses the pqueries func

jwt
