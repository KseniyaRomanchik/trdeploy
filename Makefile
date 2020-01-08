
build:
	bash -c "sudo docker run -it --rm -v $(PWD)\:/go/src/trdeploy golang\:1.13.5-alpine3.11 /bin/sh -c 'cd /go/src/trdeploy && mkdir -p cmd && go build -o ./cmd'"