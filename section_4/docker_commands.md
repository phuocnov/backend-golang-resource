** DOCKER COMMAND:

1. List all running containers

``` bash
docker ps
```
use -a to list all containers

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
6. Stop container logs 

``` bash
docker stop <container_name>
```

7. start a container

``` bash
docker start <container_name> 
```
use -a to list all containers
*** Refrences:

1. Run docker images (postgres as example): https://hub.docker.com/_/postgres

Container enviroment that run images (linux)

