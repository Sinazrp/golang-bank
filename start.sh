#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DATABASE_URL" -verbose up

echo "start server"
exec /app/main