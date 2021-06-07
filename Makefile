#makefile
GO := go
NAME := oilio
VERSION := 1.0.0
DIST := $(NAME)-$(VERSION)

all: test build

setup:

test: setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./...)

build: setup
	$(GO) build -o $(NAME) main.go

clean:
	$(GO) clean
	rm -rf $(NAME) dist