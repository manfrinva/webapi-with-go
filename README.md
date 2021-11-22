# WEBAPI with GOLANG

### API Docs

- JSON file at docs/insomnia_file.json

### Run Database

```
$ docker-compose up --build
```

### Run Main

```
$ go run main.go
```

### Migrations

- Migrations is auto runned when start the database

### Seeds

- Run seeds is manual, you have to run

```
$ go run database/seeder/main.go
```
