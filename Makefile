IMAGES = simple_ms-monolit_app

MIRROR = postgres monolit_app nginx

all:

clean:
	docker rm $(MIRROR)
	docker rmi $(IMAGES)

up:
	docker compose up

down:
	docker compose down