# Atama API
Atama API project contains implementations for standalone HTTP web server and AWS Lambda functions.

## Functions
- [ ] Get pair configuration
- [x] Get all pairs
- [x] Get pairs
- [x] Get meetings
- [x] Get goals
- [ ] Get pair selection list
- [ ] Create pair
- [ ] Get received pair request list
- [ ] Approve pair request
- [ ] Reject pair request list
- [ ] Get sent pair request list

## Structure
Project has different layers for the sake of modularity. Since that it supports both web server and AWS Lambda, 
they use common modules that contain the logic.

```
Request -> Server -> HTTP Handler  
                                  \
                                    -> Controller -> Service
                                  /
Request -> AWS Lambda Function --
```

### Web Server and AWS Lambda handlers
These are two interfaces that the API is exposed to. They are responsible of parsing the request,
building and passing the request parameters to the controllers and building the response.

### Controller
Gets the request data from Lambda or web server handlers, executes the logic through services and returns the result.

### Service
Services are responsible of data manipulation and it's usually used with databases.

## Integration Tests

Integration test suite connects to the Mongo Database running on `localhost` without any username and password.
Initially, it clears the database and inserts the sample data into the `dvn-atama-api-integration-test` database from the files in `integrationtests/data`.

```
go test ./integrationtests/...
```

## Deployment

### AWS Lambda with Serverless Framework

This commands executes the command in `Makefile` which clears the previous builds,
generates executables and deploys the AWS Functions through Serverless Framework.
```
make deploy-mentornity
```
