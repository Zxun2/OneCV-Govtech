# README

## Common docker commands

```
docker pull mysql:8.0 # pull the docker image
docker run --name mysql-root -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.0 # run the mysql database
docker exec -it mysql-root mysql -u root -p # password: secret
docker stop mysql-root # stop the container
```
