#!/bin/bash

# Ensure PostgreSQL is running via Docker
docker compose up -d

# Run goku with proper DB settings
DB_HOST=localhost DB_PORT=5433 DB_PASSWORD=postgres ./goku "$@"