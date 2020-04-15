#!/usr/bin/make

.DEFAULT_GOAL := help



help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo "\n  Allowed for overriding next properties:\n\n\
		Usage example:\n\
	    	make run"

backend-up:
	cd ./backend && docker-compose up --build -d
frontend-dep:
	cd ./frontend && npm install
frontend-up:
	cd ./frontend && npm run build
postgres-up:
	cd ./postgres && docker-compose up --build -d
mongo-up:
	cd ./mongo && docker-compose up --build -d

backend-down:
	cd ./backend && docker-compose down
postgres-down:
	cd ./postgres && docker-compose down
mongo-down:
	cd ./mongo && docker-compose down

up: postgres-up frontend-dep frontend-up mongo-up backend-up

down: backend-down postgres-down mongo-down