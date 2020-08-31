#!/usr/bin/make

include .env
export

.DEFAULT_GOAL := help

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo "\n  Allowed for overriding next properties:\n\n\
		Usage example:\n\
	    	make run"

f=cover.out

build: ## build all containers (docker compose)
	docker-compose build

up: ## build & start the project (docker-compose)
	docker-compose up --build -d

dev:
	cd ./src/frontend && npm run serve

up-dev: up front-dev

git-prep: backend-fmt backend-vendor front-build
	git add * .env.example .env.release

backend-fmt:
	cd ./src/backend && gofmt -w -s .

backend-vendor:
	cd ./src/backend && go mod vendor

git-minor:
	git add * ; git commit -m "minor" ; git push


front-build:
	cd ./src/pwa-frontend && npm run build

front-dep:
	cd ./src/pwa-frontend && npm install

front-dev:
	cd ./src/pwa-frontend && npm run serve

back-fmt:
	cd ./src/backend && gofmt -w -s .

pause:
	docker-compose pause

unpause:
	docker-compose unpause

down: ## stop the project (docker-compose)
	docker-compose down
