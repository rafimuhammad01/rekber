test:
	go test -coverprofile cover.out -v ./...

migrate-up:
	migrate -database postgres://postgres:rekberapp@localhost:5432/postgres?sslmode=disable -path postgres/migrations up $(count)

migrate-down:
	migrate -database postgres://postgres:rekberapp@localhost:5432/postgres?sslmode=disable -path postgres/migrations down $(count)

migrate-create : 
	migrate create -ext sql -dir postgres/migrations -seq $(name)