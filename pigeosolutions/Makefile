TAG=latest
IMAGE=pg_featureserv

all: docker-build docker-push

docker-build-origin:
	docker build -f Dockerfile -t pramsey/pg_featureserv:latest .

docker-build:
	docker build -f pigeosolutions/Dockerfile.pigeo -t pigeosolutions/${IMAGE}:${TAG} ./pigeosolutions

docker-push:
	docker push pigeosolutions/${IMAGE}:${TAG}
