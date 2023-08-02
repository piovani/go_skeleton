#!/bin/bash
help:
	@echo "Go Skeleton"
	@echo "https://github.com/piovani/go_skeleton"
	@echo "-----------------------------------------------"
	@echo "COMMANDS:                                      "
	@echo "make help     # prints usage info              "
	@echo "make start    # start Rest API                 "

start:
	go run main.go api

mock:
	~/go/bin/mockgen --source=entites/example/contract.go --destination=infra/test/mock/example.go -package=mock

cover:
	go test ./... -coverprofile=covarage.out -covermode=count
	go tool cover -func=covarage.out