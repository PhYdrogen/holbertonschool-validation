build: ## compile the source code
	make lint && go build
	
run: ## Run the application in background
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop the application with the 
	pgrep awesome-api | xargs kill

clean: ## Stop the application
	-make stop
	-rm ./awesome-api && rm ./awesome-api.log
	-rm coverage-integrations.out
	-rm coverage-units.out

test: ## test the server
	make unit-tests
	make integration-tests

lint: ## check if the code is correct
	golangci-lint run

unit-tests: ## run list of unit test
	go test -v -short -coverprofile=coverage-units.out

integration-tests: ## run list of integration tests
	go test -v -coverprofile=coverage-integrations.out
help: ## Display this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'