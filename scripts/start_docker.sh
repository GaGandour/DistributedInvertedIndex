cd ..
docker-compose -f docker-compose-$1.yml up -d && docker attach dii-master
docker stop $(docker ps -a -q)