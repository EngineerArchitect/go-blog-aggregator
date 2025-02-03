# Blog Aggregator


## Configuratin Database

login to postgre

```shell
sudo -u postgres psql
```

Create Database for this project
```sql
CREATE DATABASE gator;
```

Connect to the new database:
```
\c gator
```

## Configing Goose

Install Goose using

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Check installed correctly by running

```
goose -version
```

run db migrations using

The connection string takes the form
```
protocol://user:pass@host:port/database
```

for example:
```
postgres://user:passd@host:5432/gator
```

test connection string using
```shell
psql "postgres://user:passd@host:5432/gator"
```

Run migrations as follows

```shell
cd  sql/schema
goose postgres "postgres://user:passd@host:5432/gator" up
```

## SQLC

Install SQLC using the following command

```shell
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Generate GO SQL using

```
sqlc generate
```