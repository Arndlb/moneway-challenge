# moneway-challenge
Moneway Company Challenge - Transaction Simulation with Balance Update

## Build and Run

#### Build project
`./build.sh`

#### Run project 
`sudo docker run -it scylladb`
`./balance`
`./transaction`

## Project architecture

The project is divided into four parts :

- Api (api directory)
- Database (scylladb directory)
- Balance service (services/balance directory)
- Transaction service (services/transaction directory)

#### Api part

This part contains protobuf files :

 - balance.proto
 - transaction.proto 
 
 ##### Generate proto files
 
`protoc --go_out=plugins=grpc:. ./api/v1/transaction.proto`
`protoc --go_out=plugins=grpc:. ./api/v1/balance.proto`
 
 #### Database
 
 The database is a scyllaDB database, the directory only contain a dockerfile in ordre to run it and a wrapper to create a keyspace "bank"
 
 ##### Build and run
 `sudo docker build  -t scylladb .`
 `sudo docker run -it scylladb` 
 
 The database is running on localhost:9042
 
 #### Balance service 
 
 This service update the balance of an account and communicate with transaction service
 
 - db directory contains models and connections function
 - proto directory contains grpc server and service functions
 
 ##### Build and run
 `go build -o balance ./services/balance/main.go`
 `./balance`

The balance service is running on localhost:8080 

 #### Transaction service 
 
 This service creates and update transaction and communicates with balance service 
 
 - db directory contains models and connections function
 - proto directory contains grpc server, client and service functions
  
  ##### Build and run
  `go build -o transaction ./services/trnasaction/main.go`
  `./balance`
  
  The transaction service is running on localhost:8081