# Telegram Bot

Start kit for Telegram bot. Collecting all messages and users from  bot chat. Logging system connected.


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
  - send message from bot
  - message saved
  - user saved
    
### Error
* telegram
  - processing message error
  - message save error
  - user save error