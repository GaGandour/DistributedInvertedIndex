docker stop $(docker ps -a -q)
docker container prune
docker-compose up -d && docker attach dii-dii-master-1