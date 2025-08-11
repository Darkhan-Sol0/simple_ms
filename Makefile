IMAGES = simple_ms-app

MIRROR = postgres app

all:

clean:
	docker rm $(MIRROR)
	docker rmi $(IMAGES)

up:
	docker compos up

down:
	docker compos down