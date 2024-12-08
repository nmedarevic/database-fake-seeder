#!/bin/bash

echo "Started running migrations"
cd ../db/migrations

../goose/goose postgres "postgresql://postgres:example@localhost:5432/seed_db?sslmode=disable" up -allow-missing

echo "Migrations finished!"