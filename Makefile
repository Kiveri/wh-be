migrate-up:
	goose -dir ./db/migrations postgres "host=localhost user=postgres dbname=postgres password=postgres sslmode=disable" up

migrate-down:
	goose -dir ./db/migrations postgres "host=localhost user=postgres dbname=postgres password=postgres sslmode=disable" down

migrate-create:
	goose -dir db/migrations create $(F) sql