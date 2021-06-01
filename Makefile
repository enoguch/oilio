#makefile
GO := go
NAME := oilio
VERSION := 1.0.0
DIST := $(NAME)-$(VERSION)

all: test build

setup: #update_version

#update_version:
#	@for i in README.md docs/content/_index.md ; do \
#		sed -e 's!Version-[0-9.]*-blue!Version-${VERSION}-blue!g' -e 's!tag/v[0-9.]*!tag/v${VERSION}!g' $$i > a ; mv a $$i; \
#	done
#		sed -e 's!Docker-ghcr.io%2Ftamada%2Fwildcat%3A[0-9.]*-green!Docker-ghcr.io%2Ftamada%2Fwildcat%3A${VERSION}-green!g' $$i > a ; mv a $$i; \
#	@sed 's/const VERSION = .*/const VERSION = "${VERSION}"/g' cmd/$(NAME)/main.go > a
#	@mv a cmd/$(NAME)/main.go
#	@sed 's/ARG version=.*/ARG version=${VERSION}/g' Dockerfile > b
#	@mv b Dockerfile
#	@echo "Replace version to \"${VERSION}\""

test: setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./...)

build: setup
	$(GO) build -o $(NAME) cmd/oilio/*.go

#docs:

#docs/public:


#dist: build docs
#	@$(call _createDist,darwin,amd64)
#	@$(call _createDist,darwin,arm64)
#	@$(call _createDist,windows,386,.exe)
#	@$(call _createDist,linux,amd64)
#	@$(call _createDist,linux,386)

clean:
	$(GO) clean
	rm -rf $(NAME) dist

#define _update_docker
#	(sed -e '$$d' Dockerfile ; echo $(1)) > a
#	mv a Dockerfile
#endef

heroku:
	@$(call _update_docker,'CMD /opt/oilio/oilio --server --port $$PORT')
	heroku container:push web
	heroku container:release web
	@$(call _update_docker,'ENTRYPOINT [ "/opt/oilio/oilio" ]')