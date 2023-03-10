include app.env

DB_URL="mysql://root:${MYSQL_ROOT_PASSWORD}@tcp/${DB_NAME}"
DB_NAME="onecvdb"

TEST_DB_URL="mysql://root:${MYSQL_ROOT_PASSWORD}@tcp/${TEST_NAME}"
TEST_NAME="testdb"

mysql: 
	docker run --name mysql-root -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} -d mysql:8.0

createdb: 
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

server:
	@echo "Starting server..."
	go build -o main .
	go run .

.PHONY: mysql createdb dropdb migrateup migrateup-1 migratedown migratedown-1 sqlc start