# image-converter

docker run --name=invoices-db -e POSTGRES_PASSWORD=<dbpassword> -p 5432:5432 -d postgres

migrate -path ./schema -database 'postgres://postgres:123456@localhost:5432/postgres?sslmode=disable' up