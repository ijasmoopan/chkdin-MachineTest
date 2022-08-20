run:
	go run main.go

status:
	goose -dir=migrations postgres postgres://postgres:ijasmoopan@localhost:5432/chkdindb?sslmode=disable status

create:
	goose -dir=migrations postgres postgres://postgres:ijasmoopan@localhost:5432/chkdindb?sslmode=disable create users sql

up:
	goose -dir=migrations postgres postgres://postgres:ijasmoopan@localhost:5432/chkdindb?sslmode=disable up

down:
	goose -dir=migrations postgres postgres://postgres:ijasmoopan@localhost:5432/chkdindb?sslmode=disable down
