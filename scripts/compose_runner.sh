#!/usr/bin/env bash

set -e
set -x

COMPOSE=$(command -v docker-compose)
${COMPOSE} down -v &&
    MY_UID="$(id -u)" MY_GID="$(id -g)" DOCKER_BUILDKIT=1 ${COMPOSE} ${1} build &&
    MY_UID="$(id -u)" MY_GID="$(id -g)" ${COMPOSE} ${1} config &&
    if MY_UID="$(id -u)" MY_GID="$(id -g)" ${COMPOSE} ${1} up ${2}; then
        printf "> ${1} succeeded"
    else
        printf "> ${1} FAILED"
        exit 1
    fi
