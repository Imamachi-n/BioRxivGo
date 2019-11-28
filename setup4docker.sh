#!/bin/bash
# Database Setup
migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5433/biorxiv?sslmode=disable down
migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5433/biorxiv?sslmode=disable up

cd batch
make dev
