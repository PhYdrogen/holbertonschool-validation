build: check ## compile the source code
	make lint && go build
	hugo

run: ## Run the application in background
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop the application with the 
	pgrep awesome-api | xargs kill

clean: ## Stop the application and clean the directory
	-make stop
	-rm ./awesome-api && rm ./awesome-api.log
	-rm coverage-integrations.out
	-rm coverage-units.out
	-rm -rf dist/

test: ## test the server
	make unit-tests
	make integration-tests

lint: ## check if the code is correct
	golangci-lint run

unit-tests: ## run list of unit test
	go test -v -short -coverprofile=coverage-units.out

integration-tests: ## run list of integration tests
	go test -v -coverprofile=coverage-integrations.out

check:
	markdownlint content/ && markdown-link-check content/posts/*.md

validate:
	-python3 /opt/W3C-Validator/w3c_validator.py ./dist/index.html

help: ## Display this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'