# BioRxivGo

## Database setup

```bash
psql
create database biorxiv;
```

```bash
# Update the database
migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5432/biorxiv?sslmode=disable up

# Rollback the database
migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5432/biorxiv?sslmode=disable down
```

### Library Candidates to create Back-end server

```bash
dep ensure
```

- [mmcdole/gofeed](https://github.com/mmcdole/gofeed)
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [lib/pq](https://github.com/lib/pq)

```bash
go get github.com/mmcdole/gofeed
```
