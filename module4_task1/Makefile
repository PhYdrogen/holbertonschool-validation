build: build-docker ## Generate the website from the markdown and configuration files in the directory dist/
	docker run --volume ${GITHUB_WORKSPACE}/module4_task1/:/go awesome:build ls && go build
	hugo > /dev/null 2>&1
clean:    ## Cleanup the content of the directoy dist/
	rm -rf ./dist/ awesome-api awesome-api.log coverage-units.out coverage-integrations.out
	make stop
post:  ## Create a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME
	hugo new posts/$(POST_NAME).md
	sed -i '/title:/c\title: "$(POST_TITLE)"' content/posts/$(POST_NAME).md
check: ## Checking markdown files
	markdownlint content/ && markdown-link-check content/posts/*.md
validate: ## validates using Holberton tool
	python3 opt/W3C-Validator/w3c_validator.py dist/index.html || true
run: ## runs the api and logs it
	./awesome-api > ./awesome-api.log 2>&1 &
stop: ## kills the process with PID number
	pkill awesome-api
lint: ## check for potential errors
	markdownlint -f DEPLOY.md README.md || true
test: ## test hitting the endpoint
	make unit-tests
	make integration-tests
	make validate
unit-tests: build-docker ## run unit-tests
	docker run --volume ${GITHUB_WORKSPACE}/module4_task1/:/go awesome:build ls && go test -v -short -coverprofile=coverage-units.out
integration-tests: build-docker ## run intregrated tests
	docker run --volume ${GITHUB_WORKSPACE}/module4_task1/:/go awesome:build ls && go test -v  -coverprofile=coverage-integrations.out
package: build ## compile everything and zipit
	zip -r awesome-website.zip dist/ awesome-api
build-docker: ## build an image docker for each rule
	docker build -t awesome:build .build
help: ## Display this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'