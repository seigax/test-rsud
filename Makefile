dep:
	go mod download

migrate-up:
	go run apps/migration/main.go up

migrate-down:
	go run apps/migration/main.go down

migrate-status:
	go run apps/migration/main.go status

run-api:
	go run apps/api/main.go