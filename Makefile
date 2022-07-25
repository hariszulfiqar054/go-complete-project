run:
	go run main.go

create-migration:
				migrate create -ext sql -dir src/migrations -seq $(migration)	

migrate-up:
		migrate -path src/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/go-store" -verbose up

migrate-down:
			migrate -path src/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/go-store" -verbose down
			