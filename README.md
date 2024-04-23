# simple-api-go
A simple RESTful API using Go, Serverless, AWS API Gateway, Lambda, and DynamoDB.

# Project specifications

# Setup
To setup project follow thies steps :
```bash 
# First install all module dependencies
go mod tidy

# Up and run Database.
# If you execute project locally, setup DynamoDB via docker. Otherwise on AWS you need to define Instance.
docker compose up
## Execute bash script to define DynamoDB table and seed data at local.
./schema/schema-seed-data.sh

```
# Folder structure

# Configuration

# Deploy on AWS

 
