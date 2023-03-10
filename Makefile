include app.env

DB_URL="mysql://root:${MYSQL_ROOT_PASSWORD}@tcp/${DB_NAME}"
DB_NAME="onecvdb"

TEST_DB_URL="mysql://root:${MYSQL_ROOT_PASSWORD}@tcp/${TEST_NAME}"
TEST_NAME="testdb"

setup: 
	@echo "Starting MySQL..."
	docker run --name mysql-root -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} -d mysql:8.0

createdb: 
	@echo "Creating main database..."
	docker exec -it mysql-root mysql -u root --password=${MYSQL_ROOT_PASSWORD} -e "create database ${DB_NAME}"

dropdb: 
	docker exec -it mysql-root mysql -u root --password=${MYSQL_ROOT_PASSWORD} -e "drop database ${DB_NAME}"

createtestdb:
	docker exec -it mysql-root mysql -u root --password=${MYSQL_ROOT_PASSWORD} -e "create database ${TEST_NAME}"

droptestdb: 
	docker exec -it mysql-root mysql -u root --password=${MYSQL_ROOT_PASSWORD} -e "drop database ${TEST_NAME}"

migrateup:
	@echo "Migrating..." 
	migrate -path db/migrations/ -database ${DB_URL} -verbose up

migrateup-1:
	@echo "Migrating..." 
	migrate -path db/migrations/ -database ${DB_URL} -verbose up 1

migratedown:
	@echo "Migrating..." 
	migrate -path db/migrations/ -database ${DB_URL} -verbose down

migratedown-1:
	@echo "Migrating..." 
	migrate -path db/migrations/ -database ${DB_URL} -verbose down 1

sqlc: 
	sqlc generate

run:
	@echo "Running server in development mode..."
	go build -o main .
	go run .

lint:
	@echo "Running formatting..."
	@go fmt ./...
	@echo "Running linter..."
	@golangci-lint run --fix

test:
	@echo "Running tests..."
	go test ./tests

clean:
	@echo "Removing build files and cached files..."
	@rm -rf ${BUILD_DIR}
	@go clean -testcache

.PHONY: setup createdb dropdb migrateup migrateup-1 migratedown migratedown-1 sqlc start lint test clean