SOURCE=Dockerfile
IMAGE=irto/api:latest
CONTAINER=api-server

build: Dockerfile
	docker image build -f ${SOURCE} -t ${IMAGE} .

shell:
	docker container run -it --rm -v ${PWD}:/go/src/work ${IMAGE}
