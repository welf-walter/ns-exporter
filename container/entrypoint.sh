#!/bin/sh
# see https://gitlab.com/DjPicLLC/docker-cron-image/-/blob/main/entrypoint.sh

echo "Starting DockerContainer with cron job ..."

crond -f -l 8 -d 8 -L /dev/stdout
