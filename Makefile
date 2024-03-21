# Bash to execute commands
.SHELL := /usr/bin/bash -e

# One shell for each target
.ONESHELL:

.SHELLFLAGS = -ec

# PHONY target 
.PHONY: build_image run_image run_docker_compose

help: # Show help
	@echo "Available commands:"
	@egrep -h '\s#\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?# "}; \
		{printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'


build_image: # build image
	@go mod vendor
	@docker build --platform="linux/amd64" -f deployments/Dockerfile -t go-test .

run_image: # Run image
	@dotenv run docker run -d -p 8080:8080 -t go-test

run_docker_compose: # Run postgres, redis
	@docker-compose -f deployments/docker-compose.yml up -d

