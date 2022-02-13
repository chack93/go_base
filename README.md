# go_base

App description

## run

run api locally

```
HOST=localhost \
PORT=8080 \
DATABASE_URL=postgresql://@localhost/db_override \
make run
```

## run docker container locally

Run docker container in an environment similar to production, locally.  
No ports are exposed, use a revers proxy.
For example Caddy: https://caddyserver.com/docs/quick-starts/reverse-proxy

```
HOST=localhost \
PORT=8080 \
DATABASE_URL=postgresql://@localhost/app_name \
make run
```

## deployment

### build & upload image

Builds a docker image for linux/amd64 & uploads it to the given GitHub Container Registry.  
Image name is APP_NAME:VERSION, check Makefile.

```
GHCR_PAT="GitHub Container Registry Private Access Token" \
GHCR_USER="GitHub Container Registry Username" \
make release
```

### deploy to server

1. Connects to REMOTE_SERVER via ssh.
2. Pulls the image APP_NAME:VERSION & :lastest from GitHub Container Registry.
3. Stop/Remove old container & start new one with :lastest image.

No ports are exposed, use a revers proxy.
For example Caddy: https://caddyserver.com/docs/quick-starts/reverse-proxy

```
REMOTE_SERVER="username@example.com" \
GHCR_PAT="GitHub Container Registry Private Access Token" \
GHCR_USER="GitHub Container Registry Username" \
DATABASE_URL=postgresql://@localhost/app_name \
make deploy
```
