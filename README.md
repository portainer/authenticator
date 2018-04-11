# Authenticator

A tool to help you populate your `config.json` with the Portainer API Authorization header.

## Usage

```
$ docker run --rm -v ~/.docker/config.json:/config.json portainer/authenticator http://PORTAINER_URL:9000 username password
```
