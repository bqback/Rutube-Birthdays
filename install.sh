#!/bin/sh

MIGRATOR_USER=$1
MIGRATOR_PASSWORD=$2
POSTGRES_USER=$3
POSTGRES_PASSWORD=$4
POSTGRES_DB=$5
POSTGRES_HOST=$6
JWT_SECRET=$7
if [ -f "config/.env" ]; then 
    echo ".env file found"
else 
    echo ".env file not found, creating"
    echo "MIGRATOR_USER=\"$2\"" >> config/.env 
    echo "MIGRATOR_PASSWORD=\"$2\"" >> config/.env 
    echo "POSTGRES_USER=\"$3\"" >> config/.env
    echo "POSTGRES_PASSWORD=\"$4\"" >> config/.env
    echo "POSTGRES_DB=\"$5\"" >> config/.env
    echo "POSTGRES_HOST=\"$6\"" >> config/.env
    echo "JWT_SECRET=\"$7\"" >> config/.env
fi

echo "Building Docker image"
docker compose up --build --detach