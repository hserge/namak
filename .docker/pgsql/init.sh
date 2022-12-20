#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$PGSQL_USER" <<-EOSQL
	CREATE USER $APP_PGSQL_USER WITH PASSWORD '$APP_PGSQL_PASSWORD';
	CREATE DATABASE $PGSQL_DATABASE;
	GRANT ALL PRIVILEGES ON DATABASE $PGSQL_DATABASE TO $APP_PGSQL_USER;

	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL

# you need to give execution permission to your script file, since Docker copies over permissions
# chmod +x .docker/pgsql/init.sh