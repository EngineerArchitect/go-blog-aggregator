# Blog Aggregator

## Installation

Postgres and Go are required to run the program.

Install the program using the `go install` command:

1. Install the program:
```shell
go install github.com/yourusername/blog-aggregator@latest
```

2. Run the program:
```shell
blog-aggregator <command>
```

## Commands

The Blog Aggregator provides several commands to interact with the system:

- `login <username>`: Sign in user.
- `register <username>`: Register a new user.
- `users`: List all registered users.
- `agg <time-between-reqs>`: Start collecting feeds at specified interval.
- `addfeed <name> <url>`: Add a new feed with the specified name and URL.
- `feeds`: List all available feeds.
- `follow <url>`: Follow a feed with the specified URL.
- `following`: List all feeds the current user is following.
- `unfollow <feed-url>`: Unfollow the feed with the specified URL.
- `browse [limit]`: Browse latest posts [with an optional limit].

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