# BioRxivGo

**This was created during my time as a student at [Code Chrysalis](https://www.codechrysalis.io/).**

The Curated Preprint Server for Genomics & BioInformatics.

## Technologies

![tech](./img/technologies.png)

## Setup environment

### 1. PostgreSQL Database

You will need postgres installed. If you haven't installed it already, download and install the [PostgresApp](https://postgresapp.com/) and verify its working by running the command `psql` in your terminal.

Create a database for this project by running:

```bash
echo "CREATE DATABASE biorxiv | psql
```

#### Database migration & Data importing

You will need `migrate` installed. If you haven't installed it already, execute `brew install golang-migrate` in your MacOS.

Then, create a database schema using the following command.

```bash
# Update the database
migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5432/biorxiv?sslmode=disable up

# Rollback the database
migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5432/biorxiv?sslmode=disable down
```

### 2. dep: Go dependency management tool

You will need `dep` installed. If you haven't installed it already, execute the following commands in your MacOS.

```bash
brew install dep
brew upgrade dep
```

### 3. Starting REST API server

```bash
cd server
make run
```

### 4. Starting Front-end server

```bash
cd client
yarn serve
```

### The list of libraries used for Back-end server

- [mmcdole/gofeed](https://github.com/mmcdole/gofeed)
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/jinzhu/gorm)
- [lib/pq](https://github.com/lib/pq)
- [testify](https://github.com/stretchr/testify)
- [dep](https://github.com/golang/dep)

## Docker Compose

### Image

![docker-compose](./img/dockerCompose.png)

### How to start each docker container

```bash
docker-compose up
./setup4docker.sh
```

### How to build each docker container image

Example:

```bash
docker build . -t biorxivgoback:dev
docker run --rm -p 9000:9000 biorxivgoback:dev
```
