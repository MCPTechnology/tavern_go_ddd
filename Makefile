SHELL:=bash

.ONESHELL: # It all runs on one single shell

MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-biultin-rules

dev_env:=$(if $(DEV_ENV), $(DEV_ENV), 'docker')
app_slug:=finance

.DEFAULT_GOAL := run # $ make runs the application by default

#* Docker Commands
docker_init:
	sudo service --status-all
	sudo service docker start

docker_kill: permission
	@scripts/docker/kill.sh

docker_prune: permission
	@scripts/docker/prune.sh

#* Bash commands
.PHONY: clear
clear:
	@clear

.PHONY: permission
permission:
	chmod -R 777 .
	
init_alembic:
	alembic init alembic

.PHONY: clean
clean_pycache:
	rm -rf __pycache__

setup:
	@scripts/setup/local_setup.sh

get_dependencies:
	go mod tidy

run:
	cd cmd
	go run main.go