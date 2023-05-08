build: ## compile the source code
	make lint && go build
	
run: ## Run the application in background
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop the application with the 
	pgrep awesome-api | xargs kill

clean: ## Stop the application
	-make stop
	rm ./awesome-api && rm ./awesome-api.log

test: ## test the server
	curl http://localhost:9999/ 
	curl http://localhost:9999/health

lint: ## check if the code is correct
	golangci-lint run

help: ## Display this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'