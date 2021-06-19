# GO-KAFKA-RESTFUL
This repository is an experiment project to create Restful API using Gin-gonic and publish+subscribe to Kafka topic itself

## Running This App
Deploy Zookeeper, Kafka, Postgres, and Adminer locally

`docker compose up -d`

Start Go server

`go run main.go`

## Insert User Data and Produce, Subscribe to Event

```
POST locahost:8080/user

{
    "firstName": "Chan",
    "lastName": "Suttichujit
}
```

## Get User Data by ID

```
GET locahost:8080/user/:id
```
