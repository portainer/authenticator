# Authenticator

A tool to help you populate your `config.json` with the Portainer API Authorization header.

## Usage

```
$ docker run --rm -v ~/.docker/config.json:/config.json portainer/authenticator PORTAINER_URL:9443 username password true
```

## Docker CLI + Portainer API

```
$ docker -H PORTAINER_URL:9000/api/endpoints/1/docker ps -a

CONTAINER ID        IMAGE                 COMMAND                  CREATED             STATUS                         PORTS                                            NAMES
04e273b9cb27        portainer/base        "/app/portainer --no…"   5 minutes ago       Up 5 minutes                   0.0.0.0:9000->9000/tcp   portainer
...
```
