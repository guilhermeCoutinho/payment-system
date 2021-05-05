#!/bin/bash

export PGCONTAINERNAME=payment-system-postgres
export PGUSER=payment-system-user
export DBNAME=payment-system
export PROJECT_NAME=payment-system

echo "=> Starting databases"
docker-compose \
  --file dev/docker-compose.yaml \
  --project-name=$PROJECT_NAME \
  up --no-recreate -d postgres

until docker exec $PGCONTAINERNAME pg_isready
  do echo "=> Waiting for Postgres..." && sleep 1
done

docker exec $PGCONTAINERNAME psql -U $PGUSER -d $DBNAME -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp"'

echo "=> Migrating DB"

pushd migrations; make migrate; popd
echo "=> Starting services"

docker-compose \
  --file dev/docker-compose.yaml \
  --project-name=$PROJECT_NAME \
  up --no-recreate -d
