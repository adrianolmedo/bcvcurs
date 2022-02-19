.PHONY: *
.DEFAULT_GOAL := docker/app/run

export DOCKER_BUILDKIT?=1

docker/app/build:
	docker build --tag adrianolmedo/bcvcurs:dev-local .

docker/app/run: docker/app/build
	docker run \
		--rm \
		-d \
		-p 8080:80 \
		--name bcvcurs \
		adrianolmedo/bcvcurs:dev-local

docker/app/stop:
	docker stop bcvcurs

docker/app/restart: docker/app/stop docker/app/run