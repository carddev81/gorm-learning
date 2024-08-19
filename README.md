# gorm-learning
Used for researching GORM features.

## Starting database
Use the docker-compose command below. Once started you can access the adminer web app hosted here http://localhost:8080/

- docker-compose --project-name gormtester -f docker-compose.yml up -d

## Tear down database and Docker containers
Use the following commands to tear down database and docker container.  Also added something for removing the volume from docker.

- docker-compose --project-name gormtester -f docker-compose.yml down
- docker volume rm gormtester_gorm_data