docker-compose up -d
docker-compose exec app make protoc-js
docker-compose exec app make protoc
docker-compose down
make copy-to-next