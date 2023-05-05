# Food Votes

Frontend and backend code for a food voting app where users can vote on their favorite food and where they want to go with their friends that night.

## Status

Build & Tests: ![Run Tests](https://github.com/awhipp/food-votes/actions/workflows/run-tests.yml/badge.svg)](https://github.com/awhipp/food-votes/actions/workflows/run-tests.yml)

## Backend API Code

* Serverless Lambda
* Redis
* Golang 1.x

### Golang Setup

* Install Golang 1.x
* Install dependencies
* Build desired function ie: `go build -o bin/main main.go` 
** Local development ensure you are not calling the lambda.Start function and instead just calling the handler directly
* To deploy to AWS run the build.ps1 script (ensure awscli is setup appropriately)

### Functions

#### Search

Gets the local restaurants from Foursquare API for a given location or zipcode

#### Join 

Joins a Room in Progress

#### Vote

Adds votes to a given voting room

## Frontend

TBD