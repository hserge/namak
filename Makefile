include .env

help:
		@echo ""
		@echo "usage: make COMMAND"
		@echo ""
		@echo "Commands:"
		@echo "  start		       	Create and start containers"
		@echo "  stop				Stop and clear all services"
		@echo "  logs				Show docker logs"
		@echo "  migrate		    Upgrades the application by applying new migrations."
		@echo "  migrate_down		Downgrades the application by reverting old migrations."
		@echo "  migrate_to		  	Upgrades or downgrades till the specified version."
		@echo "  migrate_create     Creates a new migration."
		@echo "  migrate_history    Displays the migration history."
		@echo "  update_composer    Update PHP dependencies with composer on phpfmp container"
		@echo "  connect_mysql      Connect to the MySQL container"
		@echo "  connect_phpfpm     Connect to the PHP-FPM container"
		@echo "  clean		       	Clean junk files"


stop:
		docker-compose down -v

start:
		docker-compose -f docker-compose.yml up -d --remove-orphans

restart: stop start

logs:
		docker-compose logs -f

app_logs:
		tail -f dot818/www/protected/runtime/application.log

clean:
		rm -Rf ./.docker/data/logs/*

# PostgreSQL
pgsql_connect:
		docker exec -it $(PGSQL_CONTAINER_NAME) sh
pgsql_db_init:
		docker-compose exec db /docker-entrypoint-initdb.d/init.sh
pgsql_db_drop_db:
		docker exec -it $(PGSQL_CONTAINER_NAME) dropdb -e --if-exists --username=$(PGSQL_USER) $(PGSQL_DATABASE)
pgsql_db_drop_user:
		docker exec -it $(PGSQL_CONTAINER_NAME) dropuser -e --if-exists --username=$(PGSQL_USER) $(APP_PGSQL_USER)
pgsql_db_drop: pgsql_db_drop_db pgsql_db_drop_user

# Migration (runs on local machine)
migrate_create: # make migrate_create name=initial_migration
		migrate -verbose -database $(PGSQL_DSN) create -ext sql -dir ./migrations $(name)
migrate:
		migrate -verbose -path ./migrations -database $(PGSQL_DSN) up
migrate_down:
		migrate -verbose -path ./migrations -database $(PGSQL_DSN) down
migrate_force_version: # make migrate_force_version version=1
		migrate -verbose -path ./migrations -database $(PGSQL_DSN) force $(version)

build:
		go build -o ./bin -v
test:
		go test -v -cover ./...