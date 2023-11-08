if [ -z "$1" ]
  then
    echo "No argument supplied."
    echo "Usage: ./start_docker.sh <number of workers>"
    exit 1
fi
python3 write_docker_compose.py $1 > ../docker-compose.yml
cd ..
docker-compose -f docker-compose.yml up -d
echo "Starting DII with $1 workers"
docker attach dii-master
docker stop $(docker ps -a -q)