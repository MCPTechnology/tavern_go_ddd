#!/bin/bash

if [ ! -f "./.env" ]; then
    cp ./env/local.env ./.env
    printf "\n.env file created successfully on root."
fi

printf "> Setting up docker containers"
eval "export $(grep -v '^x' .env | xargs)"

scripts/compose_runner.sh '-f docker-compose.yaml' -d

printf "\n > Local setup is finished.\n\n"
