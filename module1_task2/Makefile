build: ## Compile the template to site prod ready
	hugo
clean: ## Clean the dist/ directory
	rm -rf dist/
post: ## Need two global variable POST_NAME and POST_TITLE
	hugo new posts/${POST_NAME}.md
	python3 script.py 
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'