# image-converter

docker run --name=postgres -e POSTGRES_PASSWORD=123456 -p 5432:5432 -d postgres

migrate -path ./schema -database 'postgres://postgres:123456@localhost:5432/postgres?sslmode=disable' up(down)

docker run -d --name redis-mq -p 6000:6379 redis
