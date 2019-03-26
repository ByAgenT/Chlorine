docker-compose up --build
docker rmi $(docker images -q -f dangling=true)