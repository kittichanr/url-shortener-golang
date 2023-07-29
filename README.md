# URL Shortener

this repo can be use in only local environment not serve into the server.

## Installation

### Redis

- Install redis on machine

```sh
brew install redis
```

- Starting and stopping Redis using launchd

```sh
// start redis
brew services start redis

// stop redis
brew services stop redis
```

- more detail other command for redis `https://redis.io/docs/getting-started/installation/install-redis-on-mac-os/`

### Docker for serve Redis and Api Url shortener to container

- on my machine I use`podman` for serve to container

- install podman and podman compose

```sh
brew install podman

brew install podman-compose
```

- start podman and compose redis and api to container in your machine

```sh
podman machine start

podman ps //check podman machine is started

podman-compose up -d // go to root file this project and run for serve to container

// after finished to run should stop podman
podman-compose down
podman machine stop
```

## Example for use api url shortener

```sh
##POST example
curl -v -X POST http://localhost:3000/api/v1 -H "Content-Type: application/json" -d '{"url":"https://www.google.com","short":"","expiry":24}'

##GET example
curl -v http://127.0.0.1:3000/OTv0FdGU8Ng
```
