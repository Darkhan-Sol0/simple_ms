NAME = chat_service

SOURCE = cmd/auth_app.go

EXEC = app

PACKAGE = github.com/gin-gonic/gin\
					github.com/jackc/pgx\
					github.com/jackc/pgx/v5\
					github.com/jackc/pgx/v5/pgxpool\
					github.com/ilyakaznacheev/cleanenv\
					github.com/joho/godotenv\
					github.com/golang-jwt/jwt/v5\

.PHONY: build run clean mod get

build:
	mkdir -p build
	go build -o build/$(EXEC) $(SOURCE)

run:
	./build/$(EXEC)

clean:
	rm -rf build

mod:
	go mod init $(NAME)

get:
	go get -u $(PACKAGE)

goose_up:
	goose -dir migrations postgres "postgresql://user:123@0.0.0.0:5430/auth_db?sslmode=disable" up

goose_down:
	goose -dir migrations postgres "postgresql://user:123@0.0.0.0:5430/auth_db?sslmode=disable" down

