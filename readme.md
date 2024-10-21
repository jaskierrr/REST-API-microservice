# project with bank cards

swagger generate server -f ./swagger.yaml --exclude-main

docker-compose up --build -d

docker-compose up -d

docker-compose ps

docker exec -it card-project sh
