## Prerequisites

download these :

go 1.15._
npm 7+
node 14._
golangci-lint

## Lifecycle

mkae lint -> do it first, check if there is commun error in the main.go code, and lint the docker file as well
make build -> compile the main in executable and check for error in markdown
make post -> create a new post with `make POST_TITLE=welcome POST_NAME="New person here" post`
make run -> launch the server
make stop -> stop the server
make test -> test the server
make clean -> clean the directory
make check -> check the md files
make validate -> verify the index.html
make unit-tests -> show all the unit test present
make integration-tests -> do all the integration tests present
make help -> show an help
make docker-tests -> launch test for the docker
make build-docker -> build the docker to make the build the same env for everyone

## Workflow

On push, pull request and at 10:30 UTC github will check for the code
Now, the make file create a docker containner to make run go commands