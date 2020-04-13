#!/usr/bin/make

include .env
export

.DEFAULT_GOAL := help

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo "\n  Allowed for overriding next properties:\n\n\
		Usage example:\n\
	    	make run"

frontend-dep: ## install/update npm packages
	cd ./src/frontend && npm install

frontend-dev: ## run dev server
	cd ./src/frontend && npm run serve

frontend-build: ## build static files into backend service
	cd ./src/frontend && npm run build

up:
	docker-compose up --build

down:
	docker-compose down