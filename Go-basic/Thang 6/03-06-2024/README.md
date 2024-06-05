Cách Run source:

1. Cài Postgres database
```
docker run \
		--name postgres-14 \
		--rm -e POSTGRES_USER=spadyprod \
		-e POSTGRES_PASSWORD=123456 \
		-p 5432:5432 -it \
		-d postgres:14
```

2. Chỉnh sửa cấu hình kết nối database
```
var conn = db.Connection{
		Host:     "localhost",
		User:     "spadyprod",
		Password: "123456",
		DBName:   "gorm_db",
		Port:     "5432",
	}
```

3. Run app
```
go run cmd/main.go
```
