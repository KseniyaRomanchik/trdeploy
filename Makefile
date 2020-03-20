include .env
export

export IMAGE_NAME=golang:1.13.5-alpine3.11
export WORKDIR=/go/src/trdeploy
export COMMIT=$(shell git rev-parse HEAD)
export DATE=$(shell date +%d-%m-%Y__%T)

build:
	@echo '*** BUILD ***'
	@docker run --env CGO_ENABLED=0 --name tdreploy-build --rm -v $(PWD)\:$(WORKDIR) $(IMAGE_NAME) /bin/sh \
	-c 'cd $(WORKDIR) && mkdir -p cmd && \
	 go build -ldflags \
	 "-X trdeploy/flags.Image=$(IMAGE_NAME) \
	 -X trdeploy/flags.Commit=$(COMMIT) \
	 -X trdeploy/flags.Time=$(DATE) \
	 -X trdeploy/commands.TerragruntConfigName=$(TERRAGRUNT_CONFIG_NAME) \
	 -X trdeploy/commands.TerraformDir=$(TERRAFORM_DIR) \
	 -X trdeploy/flags.ConfigFileName=$(CONFIG_FILE_NAME)" \
 	 -o ./cmd'

# 	@echo '*** BUILD ***' && \
# 	@docker build . --tag trdeploy-build --build-arg IMAGE_NAME=$IMAGE_NAME --build-arg WORKDIR=$WORKDIR \
# 	@echo '*** RUN ***' && \
# 	@docker run -v $(PWD)\:$(WORKDIR) trdeploy-build

install: build
	@bash -c "echo '*** delete ~/bin/trdeploy' ; rm -fr ~/bin/trdeploy "
	@bash -c "echo '*** cp new trdeploy ';  cp cmd/trdeploy ~/bin/trdeploy"

install_global: build
	@bash -c "echo '*** delete /usr/bin/trdeploy' ; rm -fr /usr/bin/trdeploy "
	@bash -c "echo '*** cp new trdeploy ';  cp cmd/trdeploy /usr/bin/trdeploy"