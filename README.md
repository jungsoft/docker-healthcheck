# docker-healthcheck

A simple binary that can be used as healthcheck in docker containers.

This repository also contains a binary file built for Alpine.

## Usage

Include the binary in your image, either by COPYing from disk or ADDing directly from this repo, then use it as HEALTHCHECK:

```Dockerfile
FROM alpine

...

HEALTHCHECK CMD /path/to/healthcheck
```

## Building

To build it for alpine, you can use the go-alpine image, for example:

```bash
docker run --rm -it golang:1.14.3-alpine
```

Then copy the go file to the running container:

```bash
docker cp healthcheck.go CONTAINER_NAME:/go/healthcheck.go
```

Build it from the container

```bash
go build healthcheck.go
```

And copy back the binary to the host

```bash
docker cp CONTAINER_NAME:/go/healthcheck healthcheck
```
