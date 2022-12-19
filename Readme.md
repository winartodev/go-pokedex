# Go Pokedex

- [Description](#description)
- [Project Structure](#project-structure)
- [Clean Architecture](#software-design-clean-architecture)
- [Database](#database)
- [Requirements](#requirements)
- [How To Run This Project](#how-to-run-this-project)

## Description
Go Pokedex is API wants to 
 + View list of pokemons
 + View detail pokemon
 + Create new pokemon
 + Update pokemon
 + Delete pokemon
 + Create new type
 + Update type
 + Delete type
 + Catch pokemon

## Project Structure
Here for the structure code 

    ├── app 
    |   # this directory is used to store main source code for this app 
    |   
    ├── build 
    |   # build directory is used to save result binary files.
    |
    ├── config                   
    |   # config directory is used to configure all the service.
    |   # ex: read env files, configuration connection apps and database etc.
    |  
    ├── deployments
    |   # directory used to save deployment files.
    |           
    ├── docs
    |   # docs directory used to save all files.
    |   # ex: api contract, api documentations, assets etc.    
    |          
    ├── entity
    |   # entitiy directory is used to declare the struct that we use
    |   # ex: Pokemon, Type, and PokemonType
    |             
    ├── enum    
    |   # will store enum data            
    |    
    ├── helper
    |   # helper directory is use to create all function that will use to help this application
    |   
    ├── middleware
    |   # middleware will filter HTTP request and validate the authentication 
    |
    ├── repository
    |   # repository directory is use to store all database handler.
    |   # ex: MySQL, PostgreSQL, MongoDB etc. and it will Querying to Create, Update, Delete,
    |   # and fetch data. here for the sample package from the repository
        |
        ├── pokemon
            ├── pokemon.go # provide communication to database and it will be use to usecase folder 
            ├── query.go # provide query operation that will use in pokemon folder
    |  
    ├── server
    |   # server directory is use to interact with user 
    |   # like this server will accept input from the user and send to usecase layer
    |
    ├── usecase
    |   # Usecase directory will use to handle business process. 
    |   # this directory will accept input from server and depens with repository layer
    |
    ├── util
    |   # util directory is use to store all utility method 
    |
    └── env.sample
    └── go.mod
    └── go.sum
    └── Makefile
    └── README.md

## Software Design (Clean Architecture)
Referrence atricle : https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html?ref=hackernoon.com

This project adopts a clean architecture software design which means it will produce a system that: 
1. Independent of Framework 

    The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.

2. Testable

    The business rules can be tested without the UI, Database, Web Server, or any other external element.

3. Independent of UI

    The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.

4. Independent of Database

    We can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.

5. Independent of external agency

    In fact your business rules simply don’t know anything at all about the outside world.

This Project using main 4 layer: 
1. [Entity](/entity/) --> set of data struct, interface and method
2. [Repository](/repository/) --> will handle to database method: fetch, put, delete, get,
3. [Usecase](/usecase/) --> handle problem logic business.
4. [Server](/server/) -->  will decide how the data present. for the case using REST API

## Database 
This project used MySQL database and there is the reason why chose MySQL database because i need only relation with another table. Why not choose mongo or postgres?. 

Postgres is like MySQL they are is RDBMS. but we can use Postgres when we had complicated case and need good performance, because postgres have many features than other DBMS. 

MongoDB is different from MySQL and PotgresSQL or other RDBMS. mongo is NoSQL or document oriented. I didn't choose Mongo because the attributes don't always change from one data to another

## Requirements
+ Go 1.17 or later
+ MySQL 8.0 or later
+ Docker

### How to Run This Project
```sh
# clone 
git clone https://github.com/winartodev/go-pokedex.git

# move to project
cd go-pokedex

# copy env.sample 
cp env.sample .env

# make sure the contents of .env file like below
APP_ENV=local
APP_URL=127.0.0.1
APP_PORT=8080

DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=pokedex
DB_USERNAME=root
DB_PASSWORD=123

# run the application (via docker compose)
make up

# below command is executed in another terminal
curl --location --request GET 'http://localhost:8080/pokedex/pokemons'

# stop application
make stop

# destroy application
make destroy
```

## How To Run Test 
```sh
# running test 
make test 

# running coverage test 
make coverage

# show coverage test 
make shwo_coverage

# show coverage on html 
make html
```