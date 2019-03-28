---
services: cosmos-db
platforms: go
author: mimig1
---

# Azure Cosmos DB: Build a Azure Cosmos DB for MongoDB API console app with Golang and the Azure portal

This sample shows how to use the Azure Cosmos DB for MongoDB API and the Go language (Golang) to create, retrieve, update, and delete a document in a collection. 

For instructions on using this quickstart, see [Azure Cosmos DB: Build a Azure Cosmos DB for MongoDB API console app with Golang and the Azure portal](https://docs.microsoft.com/en-us/azure/cosmos-db/create-mongodb-golang).

To run this sample, you will need to:

1. Follow the instructions in the link above to create a Cosmos DB
instance that uses the MongoDB API model. 
1. Set the AZURE_DATABASE environment variable to the name of the 
Cosmos DB instance that you created
1. Set the AZURE_DATABASE_PASSWORD environment variable to the 
primary password that you can get from the Azure Portal

To build this sample, you will need to:

1. Get the dependencies by running `go get`
1. Build the sample by running `go build`

Many thanks to [Durgaprasad Budhwani](https://medium.com/@durgaprasadbudhwani) for providing the code and documentation for this sample.
