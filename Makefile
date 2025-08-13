IMAGES = simple_ms-app

MIRROR = postgres app

all:

clean:
	docker rm $(MIRROR)
	docker rmi $(IMAGES)

up:
	docker compose up

down:
	docker compose down