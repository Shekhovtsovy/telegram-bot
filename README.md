# Telegram Bot

Start kit for Telegram bot. Collecting all messages and users from  bot chat. Logging system connected. Few commands
implemented as examples.


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
* bot
  - got message from telegram
  - send message from bot
  - message saved
  - user saved
    
### Error
* default
  - can`t connect to database
  - can`t start web server
* bot
  - processing message error
  - message save error
  - user save error