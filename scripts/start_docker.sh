cd ..
docker-compose up -d && docker attach dii-dii-master-1
docker stop $(docker ps -a -q)