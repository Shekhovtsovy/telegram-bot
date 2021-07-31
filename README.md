# Telegram Bot

## Setup
* Copy `docker-compose.override.yml.dist` to `docker-compose.override.yml`
```
$ cp docker-compose.override.yml.dist docker-compose.override.yml
```
* Copy `.env.axample` в `.env`
```
$ cp .env.example .env
```
* Create a network `bot` for docker
```
$ docker network create bot
```
* Run containers
```
$ docker-compose up -d
```
* Run migrations
```
$ migrate -source file://internal/db/postgres/migrations -database postgres://admin:admin@localhost:5432/bot?sslmode=disable up
```
## Makefile
* Build application
```
$ make build
```
* Run service
```
$ make run
```

## Logs

### Level
* facility
  - message

### Info
* default
    - bot started
    - bot shutdown
* telegram
  - got message from telegram
* message
    - message saved
* user
    - user saved
### Error
* message
  - message save error
* user
  - user save error
