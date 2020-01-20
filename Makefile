
build:
	@bash -c "echo '*** BUILD ***' ; docker run --env CGO_ENABLED=0 -ti --name tddeploy-build --rm -v $(PWD)\:/go/src/trdeploy golang\:1.13.5-alpine3.11 /bin/sh -c 'cd /go/src/trdeploy && mkdir -p cmd && go build -o ./cmd '"

install: build
	@bash -c "echo '*** delete ~/bin/trdeploy' ; rm -fr ~/bin/trdeploy "
	@bash -c "echo '*** cp new trdeploy ';  cp cmd/trdeploy ~/bin/trdeploy"

install_global: build
	@bash -c "echo '*** delete /usr/bin/trdeploy' ; rm -fr /usr/bin/trdeploy "
	@bash -c "echo '*** cp new trdeploy ';  cp cmd/trdeploy /usr/bin/trdeploy"