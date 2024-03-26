# Live reload + Gin (Api) + Gorm (ORM) + Testing

# Run
```bash
$ docker compose up -d
```

# Run test with coverage
```bash
$ docker compose exec go go test -coverprofile=coverage.out ./...
```

# Generate HTML from coverage
```bash
$ docker compose exec go go tool cover -html=/app/coverage.out -o=cover.html
```