SHELL=bash

setup:
	@envsubst < ./.env.template > .env
	@cat .env

up:
	make setup
	docker-compose up -d --build

build:
	go build -o ./dist/round-robin-time-table ./main.go