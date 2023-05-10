build: check ## compile the source code 
	make update-hugo
	make go-build
	hugo

go-build: ## build the api
	go build

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

test:  ## test the 
	make unit-tests
	make integration-tests

lint: ## check if the code is correct
	golangci-lint run

unit-tests: ## run list of unit test
	go test -v -short -coverprofile=coverage-units.out

integration-tests: ## run list of integration tests
	go test -v -coverprofile=coverage-integrations.out

check: ## check the markdown to be more effective
	markdownlint content/ && markdown-link-check content/posts/*.md

validate: ## useless but can improve code
	make build
	-python3 /opt/W3C-Validator/w3c_validator.py ./dist/index.html

post: ## Need two global variable POST_NAME and POST_TITLE
	hugo new posts/${POST_NAME}.md
	python3 script.py

update-hugo: ## update hugo
	curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
	apt install ./hugo.deb
	rm hugo.deb
help: ## Display this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
	