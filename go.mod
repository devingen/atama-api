module github.com/devingen/atama-api

go 1.12

replace github.com/devingen/api-core => ../api-core

replace github.com/devingen/data-api => ../data-api

require (
	github.com/aws/aws-lambda-go v1.16.0
	github.com/devingen/api-core v0.0.0-00010101000000-000000000000
	github.com/devingen/data-api v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.7.4
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.3.2
)
