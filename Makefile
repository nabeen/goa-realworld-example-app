APP_NAME := goa-realworld-example-app
REPO := github.com/nabeen/goa-realworld-example-app

goagen:
	@goa gen $(REPO)/design

regen:
	@rm -rf cmd
	@goa example $(REPO)/design

run:
	@cd cmd/api && go run main.go http.go

clean:
	@rm -rf cmd
	@rm -rf gen

read-env:
	@. .env

# docker用 Makefile
docker-all: docker-rm docker-build docker-up docker-ps

docker-build:
	@docker-compose build

docker-up:
	@docker-compose up -d

docker-ps:
	@docker ps -a && echo "\n"
	@docker-compose ps

docker-rm:
	@docker-compose stop
	@docker-compose rm -f
