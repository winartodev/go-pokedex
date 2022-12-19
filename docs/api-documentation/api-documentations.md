# POKEDEX API DOCUMENTATION  

[<< BACK](Readme.md)

Pokedex API: 
- [Default](#default)
  - [Login](#login)
    - [Resource URL](#resource-url)
    - [Parameters](#parameters)
    - [POST Request Data](#post-request-data)
    - [Example Request](#example-request)
    - [Example Response](#example-response)
  - [Register](#register)
    - [Resource URL](#resource-url-1)
    - [Parameters](#parameters-1)
    - [POST Request Data](#post-request-data-1)
    - [Example Request](#example-request-1)
    - [Example Response](#example-response-1)
  - [Logout](#logout)
    - [Resource URL](#resource-url-2)
    - [Parameters](#parameters-2)
    - [POST Request Data](#post-request-data-2)
    - [Example Request](#example-request-2)
    - [Example Response](#example-response-2)
  - [Healhz](#healthz)
    - [Resource URL](#resource-url-3)
    - [Example Request](#example-request-3)
    - [Example Response](#example-response-3)
- [Public API](#public-api)
  - [List Of Pokemon](#list-of-pokemon)
    - [Resource URL](#resource-url-4)
    - [Parameters](#parameters-3)
    - [Example Request](#example-request-4)
    - [Example Response](#example-response-4)
  - [Detail Pokemon](#detail-pokemon)
    - [Resource URL](#resource-url-5)
    - [Parameters](#parameters-4)
    - [Example Request](#example-request-5)
    - [Example Response](#example-response-5)
  - [List Of Types](#list-of-type)
    - [Resource URL](#resource-url-6)
    - [Parameters](#parameters-5)
    - [Example Request](#example-request-6)
    - [Example Response](#example-response-6)
- [Internal API](#internal-api)
  - [List Of Pokemon](#list-of-pokemon-1)
    - [Resource URL](#resource-url-7)
    - [Parameters](#parameters-6)
    - [Example Request](#example-request-7)
    - [Example Response](#example-response-7)
  - [Create New Pokemon](#create-pokemon)
    - [Resource URL](#resource-url-8)
    - [Parameters](#parameters-7)
    - [POST Request Data](#post-request-data-3)
    - [Example Request](#example-request-8)
    - [Example Response](#example-response-8)
  - [Detail Pokemon](#detail-pokemon-1)
    - [Resource URL](#resource-url-9)
    - [Parameters](#parameters-8)
    - [Example Request](#example-request-9)
    - [Example Response](#example-response-9)
  - [Update Pokemon](#update-pokemon)
    - [Resource URL](#resource-url-10)
    - [Parameters](#parameters-9)
    - [PUT Request Data](#put-request-data)
    - [Example Request](#example-request-10)
    - [Example Response](#example-response-10)
  - [Delete Pokemon](#delete-pokemon)
    - [Resource URL](#resource-url-11)
    - [Parameters](#parameters-10)
    - [Example Request](#example-request-11)
    - [Example Response](#example-response-11)
  - [List Of Types](#list-of-type-1)
    - [Resource URL](#resource-url-12)
    - [Parameters](#parameters-11)
    - [Example Request](#example-request-12)
    - [Example Response](#example-response-12)
  - [Detail Of Types](#detail-of-type)
    - [Resource URL](#resource-url-13)
    - [Parameters](#parameters-12)
    - [Example Request](#example-request-13)
    - [Example Response](#example-response-13)
  - [Create New Types](#create-new-type)
    - [Resource URL](#resource-url-14)
    - [Parameters](#parameters-13)
    - [POST Request Data](#post-request-data-4)
    - [Example Request](#example-request-14)
    - [Example Response](#example-response-14)
  - [Update Type](#update-type)
    - [Resource URL](#resource-url-15)
    - [Parameters](#parameters-14)
    - [PUT Request Data](#put-request-data-1)
    - [Example Request](#example-request-15)
    - [Example Response](#example-response-15)
- [UserAPI](#user)
  - [Catch Pokemon](#catch-pokemon)
    - [Resource URL](#resource-url-16)
    - [Parameters](#parameters-15)
    - [POST Request Data](#post-request-data-5)
    - [Example Request](#example-request-16)
    - [Example Response](#example-response-16)

## Default
---
Default API is use to user login, register, logout, and check health api, this common used for all API

### Login 
Login uses username and password for authentication if request is valid, it will generate JWT token and saved cookie and return success message
 + use `POST` method

#### Resource URL
+ http://127.0.0.1:8080/login

#### Parameters 
None

#### POST Request Data
+ `username` *(Required)* user username
+ `password` *(Required)* user password

#### Example Request
```sh
curl -X 'POST' \
'http://127.0.0.1:8080/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "username": "username",
  "password": "123"
}'
```

#### Example Response
```json
{
  "status": 200,
  "message": "login success",
  "data": null
}
```

### Register
Register is used to register new user when register success it will return last insert id and success message 
+ use `POST` method

#### Resource URL
+ http://127.0.0.1:8080/register

#### Parameters 
None

#### POST Request Data
+ `username` *(Required)* user username
+ `email` *(Required)* user email
+ `password` *(Required)* user password
+ `role` *(Required)* user role to mapping as an admin value `1` or user value `2`

#### Example Request 
```sh
curl -X 'POST' \
  'http://127.0.0.1:8080/register' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "username": "username",
  "email": "user@mail.com",
  "password": "123",
  "role": 1
}'
```

#### Example Response
```json
{
  "status": 200,
  "message": "success create account",
  "data": 1
}
```

### Logout
Logout will delete token 
+ use `POST` method

#### Resource URL
+ http://127.0.0.1:8080/register

#### Parameters
None

#### POST Request Data
None

#### Example Request 
```sh
curl -X 'POST' \
  'http://127.0.0.1:8080/logout' \
  -H 'accept: application/json' \
  -d ''
```

#### Example Response
```json
{
  "status": 200,
  "message": "user logout success",
  "data": null
}
```

### Healthz
Healthz check if api has working or not 
+ use `GET` method

#### Resource URL
+ http://127.0.0.1:8080/healthz

#### Example Request 
```sh
curl -X 'GET' \
  'http://127.0.0.1:8080/healthz' \
  -H 'accept: text/plain'
```

#### Example Response 
```
ok
```

## Public 
---
this public API is not require authentication, this API allow user or guest to 
+ View a list of pokemon and each profile detail with an
image of it
+ Search pokemon by their name, filter by type
+ Sort pokemon by name, id and order them by ascending
or descending

### List Of Pokemon
Get number of pokemons, if parameter `name` exist it will give number of pokemon matching with name `name`. 

+ Use `GET` method

#### Resource URL
+ http://127.0.0.1:8080/pokedex/pokemons. show number of pokemons
+ http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour. show number of pokemon match with `name` 
+ http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour&type=1%2C2%2C3. show number of pokemons filter by `name` and `type`
+ http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour&options=1. show number of pokemons filter by `name` and `options`
+ http://127.0.0.1:8080/pokedex/pokemons?sort_by=id&order_by=asc. show number of pokemons with query `sort_by` and `order_by`

#### Parameters
+ `name` *(optional)*. Name use to search pokemon 
+ `options` *(optional)* Options to filter pokemon already catched or not catched. if want filter pokemon already catched use `1` and to filter pokemon has't catched use `0`
+ `type` *(optional)* Type to filter pokemon by type example value `1` to filter pokemon type Fire, or we can use multiple value to filter pokemon type. Allowed values `1,2,3` 
+ `sort_by` & `order_by` *(optional)* Sort by and Order by to sort pokemon by `id` and `name` and order by ascending or descending

#### Example Request 
```sh
curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons' \
  -H 'accept: application/json'

curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour' \
  -H 'accept: application/json'

curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour&type=1%2C2%2C3' \
  -H 'accept: application/json'

curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour&options=1' \
  -H 'accept: application/json'

curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour&options=1' \
  -H 'accept: application/json'
```

#### Example Response
```json 
{
  "status": 200,
  "message": "",
  "data": [
    {
      "id": 1,
      "name": "Wigglytuff",
      "species": "Balloon Pokémon",
      "types": [
        "NORMAL"
      ],
      "catched": 0,
      "image_url": "https://img.pokemondb.net/artwork/large/wigglytuff.jpg"
    },
    {
     
      "id": 2,
      "name": "Bulbasaur",
      "species": "Seed Pokémon",
      "types": [
        "NORMAL",
        "POISON"
      ],
      "catched": 0,
      "image_url": "https://img.pokemondb.net/artwork/avif/bulbasaur.avif"
    }
  ]
}
```

### Detail Pokemon
Get Detail Pokemon

+ use `GET` method

#### Resource URL
+ http://127.0.0.1:8080/pokedex/pokemons/:id

#### Parameters
+ `id` *(required)*. Identifier for pokemon want to show detail pokemon.

#### Example Request 
```sh
curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons/1' \
  -H 'accept: application/json'
```

#### Example Response
```json
{
  "status": 200,
  "message": "",
  "data": {
    "id": 1,
    "name": "Wigglytuff",
    "species": "Balloon Pokémon",
    "types": [
      "NORMAL"
    ],
    "catched": 0,
    "image_url": "https://img.pokemondb.net/artwork/large/wigglytuff.jpg",
    "description": "Wigglytuff is a Normal/Fairy type Pokémon introduced in Generation 1. It is known as the Balloon Pokémon.",
    "weight": 12,
    "height": 1,
    "stats": {
      "hp": 140,
      "attack": 70,
      "def": 45,
      "speed": 45
    }
  }
}
```

### List Of Type 
Show All Type of Pokemon

+ use `GET` method

#### Resource URL
+ http://127.0.0.1:8080/pokedex/types

#### Parameters
None

#### Example Request 
```sh
curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/types' \
  -H 'accept: application/json'
```

#### Example Response
```json
{
  "status": 200,
  "message": "",
  "data": [
    {
      "id": 1,
      "name": "GRASS"
    },
    {
      "id": 2,
      "name": "PSYCHIC"
    },
    {
      "id": 3,
      "name": "FIRE"
    }
  ]
}
```

## Internal API
Used for admin role, required `token` save as Cookie in header 
to validate expired time and role the user (as admin). if match user can access this path or if not match user will get 401 unauthorize. 

This API allowed admin to  list, add, update and delete pokemon

### List Of Pokemon
Get number of pokemons, if parameter `name` exist it will give number of pokemon matching with name `name`. 

+ Use `GET` method
+ Required authentication

#### Resource URL
+ http://127.0.0.1:8080/pokedex/pokemons. show number of pokemons
+ http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour. show number of pokemon match with `name` 
+ http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour&type=1%2C2%2C3. show number of pokemons filter by `name` and `type`
+ http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour&options=1. show number of pokemons filter by `name` and `options`
+ http://127.0.0.1:8080/pokedex/pokemons?sort_by=id&order_by=asc. show number of pokemons with query `sort_by` and `order_by`

#### Parameters
+ `name` *(optional)*. Name use to search pokemon 
+ `options` *(optional)* Options to filter pokemon already catched or not catched. if want filter pokemon already catched use `1` and to filter pokemon has't catched use `0`
+ `type` *(optional)* Type to filter pokemon by type example value `1` to filter pokemon type Fire, or we can use multiple value to filter pokemon type. Allowed values `1,2,3` 
+ `sort_by` & `order_by` *(optional)* Sort by and Order by to sort pokemon by `id` and `name` and order by ascending or descending

#### Example Request 
```sh
curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons' \
  -H 'accept: application/json' \
  -H 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBtYWlsIiwicm9sZSI6MiwiZXhwIjoxNjcxMjU3MTE2fQ.x0CMiefx6oEKPL5_BBQLNQblFkfrJ0yVEPjxLWG4IRQ'

curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour' \
  -H 'accept: application/json'
  -H 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBtYWlsIiwicm9sZSI6MiwiZXhwIjoxNjcxMjU3MTE2fQ.x0CMiefx6oEKPL5_BBQLNQblFkfrJ0yVEPjxLWG4IRQ'

curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour&type=1%2C2%2C3' \
  -H 'accept: application/json'
  -H 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBtYWlsIiwicm9sZSI6MiwiZXhwIjoxNjcxMjU3MTE2fQ.x0CMiefx6oEKPL5_BBQLNQblFkfrJ0yVEPjxLWG4IRQ'

curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour&options=1' \
  -H 'accept: application/json'
  -H 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBtYWlsIiwicm9sZSI6MiwiZXhwIjoxNjcxMjU3MTE2fQ.x0CMiefx6oEKPL5_BBQLNQblFkfrJ0yVEPjxLWG4IRQ'

curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/pokemons?name=Bulbasour&options=1' \
  -H 'accept: application/json'
  -H 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBtYWlsIiwicm9sZSI6MiwiZXhwIjoxNjcxMjU3MTE2fQ.x0CMiefx6oEKPL5_BBQLNQblFkfrJ0yVEPjxLWG4IRQ'
```

#### Example Response
```json 
{
  "status": 200,
  "message": "",
  "data": [
    {
      "id": 1,
      "name": "Wigglytuff",
      "species": "Balloon Pokémon",
      "types": [
        "NORMAL"
      ],
      "catched": 0,
      "image_url": "https://img.pokemondb.net/artwork/large/wigglytuff.jpg"
    },
    {
      "id": 2,
      "name": "Bulbasaur",
      "species": "Seed Pokémon",
      "types": [
        "NORMAL",
        "POISON"
      ],
      "catched": 0,
      "image_url": "https://img.pokemondb.net/artwork/avif/bulbasaur.avif"
    }
  ]
}
```

### Create Pokemon
Create new pokemon
+ Use `POST` method
+ Required authentication

#### Resource URL
+ http://127.0.0.1:8080/internal/pokedex/pokemons

#### Parameters
None

#### POST Request Data
+ `name` *(required)* Pokemon name
+ `species` *(required)* Pokemon species
+ `types` *(required)* Pokemon type
+ `catched` *(required)* Pokemon status is catched or not
+ `image_url` *(required)* Pokemon image
+ `description` *(optional)* Pokemon description
+ `weight` *(optional)* Pokemon weight
+ `height` *(optional)* Pokemon height
+ `stats` *(required)* Pokemon statistic
  + `hp` *(optional)* Pokemon HP
  + `attack` *(optional)* Pokemon attack strength
  + `deff` *(optional)* Strength of pokemon deffence 
  + `speed` *(optional)* Spped of Pokemon

#### Example Request 
```sh
curl -X 'POST' \
  'http://127.0.0.1:8080/internal/pokedex/pokemons' \
  -H 'accept: application/json' \
  -H 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBtYWlsIiwicm9sZSI6MiwiZXhwIjoxNjcxMjU3MTE2fQ.x0CMiefx6oEKPL5_BBQLNQblFkfrJ0yVEPjxLWG4IRQ' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Bulbasour",
  "species": "Fyling Pokemon",
  "types": [
    1,
    2
  ],
  "catched": 0,
  "image_url": "https://pokedex.photos/200/300",
  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
  "weight": 2.5,
  "height": 5.1,
  "stats": {
    "hp": 100,
    "attack": 100,
    "def": 100,
    "speed": 100
  }
}'
```

#### Example Response
```json 
{
  "status": 200,
  "message": "success create pokemon",
  "data": 1
}
```

### Detail Pokemon 
Get Detail Pokemon

+ use `GET` method
+ + Required authentication

#### Resource URL
+ http://127.0.0.1:8080/internal/pokedex/pokemons/:id

#### Parameters
+ `id` *(required)*. Identifier for pokemon want to show detail pokemon.

#### Example Request 
```sh
curl -X 'GET' \
  'http://127.0.0.1:8080/internal/pokedex/pokemons/1' \
  -H 'accept: application/json'
```

#### Example Response
```json
{
  "status": 200,
  "message": "",
  "data": {
    "id": 0,
    "name": "Wigglytuff",
    "species": "Balloon Pokémon",
    "types": [
      "NORMAL"
    ],
    "catched": 0,
    "image_url": "https://img.pokemondb.net/artwork/large/wigglytuff.jpg",
    "description": "Wigglytuff is a Normal/Fairy type Pokémon introduced in Generation 1. It is known as the Balloon Pokémon.",
    "weight": 12,
    "height": 1,
    "stats": {
      "hp": 140,
      "attack": 70,
      "def": 45,
      "speed": 45
    }
  }
}
```

### Update Pokemon
Update existing data pokemon if success will return updated new data

+ Use `PUT` method
+ Required authentication

#### Resource URL 
http://127.0.0.1:8080/internal/pokedex/pokemons/:id

#### Parameters
+ `id` *(required)*. Identifier for pokemon want to get specific pokemon.

#### PUT Request Data
+ `name` *(required)* Pokemon name
+ `species` *(required)* Pokemon species
+ `types` *(required)* Pokemon type
+ `catched` *(required)* Pokemon status is catched or not
+ `image_url` *(required)* Pokemon image
+ `description` *(optional)* Pokemon description
+ `weight` *(optional)* Pokemon weight
+ `height` *(optional)* Pokemon height
+ `stats` *(required)* Pokemon statistic
  + `hp` *(optional)* Pokemon HP
  + `attack` *(optional)* Pokemon attack strength
  + `deff` *(optional)* Strength of pokemon deffence 
  + `speed` *(optional)* Spped of Pokemon

#### Example Request 
```sh
curl -X 'PUT' \
  'http://127.0.0.1:8080/internal/pokedex/pokemons/1' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Bulbasour",
  "species": "Fyling Pokemon",
  "types": [
    1,
    2
  ],
  "catched": 0,
  "image_url": "https://pokedex.photos/200/300",
  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
  "weight": 2.5,
  "height": 5.1,
  "stats": {
    "hp": 100,
    "attack": 100,
    "def": 100,
    "speed": 100
  }
}'
```

#### Example Response
```json
{
  "status": 200,
  "message": "success update pokemon",
  "data": {
    "id": 0,
    "name": "Bulbasour",
    "species": "Fyling Pokemon",
    "types": [
      "NORMAL",
      "POISON"
    ],
    "catched": 0,
    "image_url": "https://pokedex.photos/200/300",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
    "weight": 2.5,
    "height": 5.1,
    "stats": {
      "hp": 100,
      "attack": 100,
      "def": 100,
      "speed": 100
    }
  }
}
```

### Delete Pokemon
Delete pokemon

+ Use `DELETE` method
+ Required authentication

#### Resource URL 
http://127.0.0.1:8080/internal/pokedex/pokemons/:id

#### Parameters
+ `id` *(required)*. Identifier for pokemon want to get specific pokemon.

#### Example Request
```sh
curl -X 'DELETE' \
  'http://127.0.0.1:8080/internal/pokedex/pokemons/1' \
  -H 'accept: application/json'
```

#### Example Response
```json
{
  "status": 200,
  "message": "delete pokemon success",
  "data": 1
}
```

### List Of Type 
Show All Type

+ use `GET` method
+ required authentication

#### Resource URL
+ http://127.0.0.1:8080/internal/pokedex/types

#### Parameters
None

#### Example Request 
```sh
curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/types' \
  -H 'accept: application/json'
```

#### Example Response
```json
{
  "status": 200,
  "message": "",
  "data": [
    {
      "id": 1,
      "name": "GRASS"
    },
    {
      "id": 2,
      "name": "Physics"
    },
    {
      "id": 3,
      "name": "Fire"
    }
  ]
}
```

### Detail Of Type 
Show Specific Type

+ use `GET` method
+ required authentication

#### Resource URL
+ http://127.0.0.1:8080/internal/pokedex/types/:id

#### Parameters
+ `id` *(required)*. Identifier for type 

#### Example Request 
```sh
curl -X 'GET' \
  'http://127.0.0.1:8080/pokedex/types/1' \
  -H 'accept: application/json'
```

#### Example Response
```json
{
  "status": 200,
  "message": "",
  "data": {
    "id": 1,
    "name": "GRASS"
  }
}
```

### Create New Type
Create new type
+ Use `POST` method
+ Required authentication

#### Resource URL
+ http://127.0.0.1:8080/internal/pokedex/types

#### Parameters
None

#### POST Request Data
+ `name` *(required)* Type name

#### Example Request 
```sh
curl -X 'POST' \
  'http://127.0.0.1:8080/internal/pokedex/pokemons' \
  -H 'accept: application/json' \
  -H 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBtYWlsIiwicm9sZSI6MiwiZXhwIjoxNjcxMjU3MTE2fQ.x0CMiefx6oEKPL5_BBQLNQblFkfrJ0yVEPjxLWG4IRQ' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Bulbasour"
  }
}'
```

#### Example Response
```json 
{
  "status": 200,
  "message": "create type success",
  "data": 1
}
```

### Update Type
Update existing data types 

+ Use `PUT` method
+ Required authentication

#### Resource URL 
http://127.0.0.1:8080/internal/pokedex/types/:id

#### Parameters
+ `id` *(required)*. Identifier for types

#### PUT Request Data 
+ `name` *(required)* types name

#### Example Request 
```sh
curl -X 'PUT' \
  'http://127.0.0.1:8080/internal/pokedex/types/1' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Bulbasour"
}'
```

#### Example Response
```json
{
  "status": 200,
  "message": "update type success",
  "data": 1
}
```

## User
---
Used for user role, required `token` save as Cookie in header 
to validate expired time and role the user (as user). if match user can access this path or if not match user will get 401 unauthorize.

### Catch Pokemon
Use to catch pokemon. update status catch to `1` catched

+ Use `POST` method

#### Resource URL
http://127.0.0.1:8080/user/pokedex/pokemons/:id/catch

#### Parameters
+ `id` *(required)*. Identifier for pokemon want to get specific pokemon.

#### POST Request Data
None

#### Expected Request
```sh
curl -X 'POST' \
  'http://127.0.0.1:8080/user/pokedex/pokemons/1/catch' \
  -H 'accept: application/json' \
  -d ''
```

#### Expected Response
```json
{
  "status": 200,
  "message": "Pokemon success catched",
  "data": 1
}
```












