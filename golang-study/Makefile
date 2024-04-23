create-migration-file:
	docker-compose exec golang-study migrate create -ext sql -dir ./database/migrations -seq ${TABLE_NAME}

migrate-up:
	docker-compose exec golang-study migrate -database "postgres://postgres:password@db:5432/golang_study?sslmode=disable" -path ./database/migrations up

migrate-down:
	docker-compose exec golang-study migrate -database "postgres://postgres:password@db:5432/golang_study?sslmode=disable" -path ./database/migrations down

migrate-force:
	docker-compose exec golang-study migrate -database "postgres://postgres:password@db:5432/golang_study?sslmode=disable" -path ./database/migrations force ${VERSION}

gen-code-from-db:
	docker-compose exec golang-study sqlboiler psql --output ./generated/database/models --pkgname models --wipe && go mod tidy