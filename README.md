# pwHasherAPI

## Installation
`cd` to root directory after cloning repository

To run: `go run .`

Golang version used: 1.18

## Testing
To run all tests: `go test -v ./...`

Test files are found within separate packages postfixed by `_test`
## Structure
```shell
- app
  Contains all app-related logic
  - launcher
    Contains App struct which encompasses the application at runtime
  - router
    Contains all routing mechanisms to setup API endpoints
  - handler
    Contains business logic functionality for each router group
  - utils
    Contains utility abstractions and implementations used throughout the app
  - model
    Plain old structs, for uniform use across the app
- main.go
  Launches the app
```

## API access
Endpoints match those specified in the document. For ease of use, a Postman JSON collection is included in the root directory of the project

## Assumptions
- A simplified in-memory datastore was implemented to keep the focus on core functionality, and not include any external dependencies. The caveat of this is that data doesn't persist between runs
- Time for processing requests under POST /hash is calculated as the time since request began until the time the hash is generated. 