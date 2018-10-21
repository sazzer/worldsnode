#!/bin/sh

echo Migrating Postgres Database
cd /migrations
db-migrate up -e dev
