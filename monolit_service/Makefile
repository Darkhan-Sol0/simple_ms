NAME = auth_service

EXEC = app

SOURCE = ./cmd/app/main.go

PACKAGE =\
					github.com/ilyakaznacheev/cleanenv\
					github.com/labstack/echo/v4\
					github.com/jackc/pgx/v5\
					github.com/jackc/pgx/v5/pgxpool\

PACKAGE_TEMPL =\
					github.com/a-h/templ/cmd/templ@latest\

PACKAGE_GOOSE =\
					github.com/pressly/goose/v3/cmd/goose@latest\

.PHONY: all build run clean init get get_install templ

all: clean build run

build:
	go build -o $(EXEC) $(SOURCE)

run:
	./$(EXEC)

clean:
	rm -f $(EXEC)

init:
	go mod init $(NAME)

get:
	go get -u $(PACKAGE)

templ:
	go tool templ generate

templ_install:
	go get -tool $(PACKAGE_TEMPL)

goose_up:
	~/go/bin/goose -dir migrations postgres "postgresql://user:password@pgdb_user:5432/db?sslmode=disable" up

goose_down:
	~/go/bin/goose -dir migrations postgres "postgresql://user:password@pgdb_user:5432/db?sslmode=disable" down

goose_install:
	go install $(PACKAGE_GOOSE)