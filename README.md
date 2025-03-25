# Webapp Template

Webapp template using golang, React and GORM (PostgreSQL).

## Init the db

mkdir -p db
initdb-17 -D db
postgres-17 -D db start  # start the server

## Run the server

cd src && (cd client && npm run build) && go run .
