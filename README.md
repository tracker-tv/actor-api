# Actor API

## Installation

You can simply run command from Makefile
```bash
make start
```

or with docker compose, will run postgresql container and execute migration from db-migration container
```bash
docker compose run --rm db-migration
```

Run the server
```bash
go run ./cmd/api
```

If you want to connect to the database from your host machine, you can use the `compose.override.yml` file
```bash
cp compose.override.yml.dist compose.override.yml
```
