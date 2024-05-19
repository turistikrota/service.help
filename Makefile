build:
	docker build --build-arg GITHUB_USER=${TR_GIT_USER} --build-arg GITHUB_TOKEN=${TR_GIT_TOKEN} -t github.com/turistikrota/service.help . 

run:
	docker service create --name help-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6029:6029 github.com/turistikrota/service.help:latest

remove:
	docker service rm help-api-turistikrota-com

stop:
	docker service scale help-api-turistikrota-com=0

start:
	docker service scale help-api-turistikrota-com=1

restart: remove build run
