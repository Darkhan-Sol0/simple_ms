NAME = myApp

EXEC = app

SOURCE = ./cmd/app/main.go

PACKAGE =\
					github.com/ilyakaznacheev/cleanenv\
					github.com/labstack/echo/v4\
					github.com/jackc/pgx/v5\
					github.com/jackc/pgx/v5/pgxpool\

PACKAGE_INST =\
					-tool github.com/a-h/templ/cmd/templ@latest\

.PHONY: all build run clean init get

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

get_install:
	go get -tool $(PACKAGE_INST)