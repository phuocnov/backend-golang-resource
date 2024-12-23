** DOCKER COMMAND:

1. List all running containers

``` bash
docker ps
```
2. List all available images


``` bash
docker images 
```

3. Pull new images

``` bash
docker pull <image_name>:<tag>
```

4. Run a command to docker container

``` bash
docker exec -it <container_name> <command (eg: psql)> [args] 
```

5. Check container logs 

``` bash
docker logs <container_name>
```

*** Refrences:

1. Run docker images (postgres as example): https://hub.docker.com/_/postgres

Container enviroment that run images (linux)

