SOURCE=Dockerfile
IMAGE=irto/frontdev:latest
CONTAINER=front-dev

# build container
build: Dockerfile
	docker image build -f ${SOURCE} -t ${IMAGE} .

# create new container and login to the shell
shell:
	docker container run -it --rm ${IMAGE}

# clean up all stopped containers
clean:
	docker container prune

# delete all image
doomsday:
	docker image rm `docker image ls -q`

elm: Main.elm
	elm make Main.elm --output=index.html
