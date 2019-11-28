#!/bin/bash
# Database Setup
migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5432/biorxiv?sslmode=disable down
migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5432/biorxiv?sslmode=disable up

cd batch
make dev
