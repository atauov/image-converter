# image-converter

docker run --name=postgres -e POSTGRES_PASSWORD=123456 -p 5432:5432 -d postgres

migrate -path ./schema -database 'postgres://postgres:123456@localhost:5432/postgres?sslmode=disable' up(down)

docker run -d --name redis-mq -p 6000:6379 redis

//TODO check original resolution of image (set min value or change other sizes) first copy > original
~~//TODO change image -~~ 
~~//TODO delete on S3~~
~~//TODO delete on local~~
//TODO change status in DB
//TODO non-public bucket
//TODO temporary key
//TODO screenshot bucket and policy
//TODO polymorphic keys in DB
//TODO clean arc, SOLID
//TODO write tests
//TODO readme
//TODO monitor bg workers
//TODO docker & docker-compose
//TODO makefile

//TODO delete from s3 when update