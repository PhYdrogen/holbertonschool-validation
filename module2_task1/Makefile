build: ## compile the source code
	golint run && go build
run: ## Run the application in background
	./awesome-api >./awesome-api.log 2>&1 &
stop: ## Stop the application with the 
	pgrep awesome-api | xargs kill

clean: stop ## Stop the application
	rm ./awesome-api && rm ./awesome-api.log 
test: ## test the server
	curl http://localhost:9999/ 
	curl http://localhost:9999/health

help: ## Display this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'