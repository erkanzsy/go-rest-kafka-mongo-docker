#!/bin/bash

setup:
	docker network create go-example
	docker-compose up --build -d

produce:
	go run producer/producer.go

consume:
	go run consumer/consumer.go