printf "Killing all Docker instances...\n"
eval "docker kill $(docker ps -q)"
