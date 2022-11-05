SHELL := /bin/bash

## Golang Stuff
GOCMD=go
GORUN=$(GOCMD) run

SERVICE=discord-spam-bot

init:
	rm -rf go.mod
	$(GOCMD) mod init $(SERVICE)

tidy:
	rm -rf go.sum
	$(GOCMD) mod tidy

run:
	ENV=local $(GORUN) main.go

build:
	$(GOCMD) build -o ./app main.go