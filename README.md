# docker-healthcheck

A simple binary that can be used as healthcheck in docker containers.

This repository also contains binary files built for Alpine AMD64 and ARM64.

## Usage

Include the binary in your image, either by COPYing from disk or ADDing directly from this repo, then use it as HEALTHCHECK:

For alpine linux AMD64:

```Dockerfile
FROM alpine

ADD https://github.com/jungsoft/docker-healthcheck/releases/download/0.1.0/healthcheck_alpine_amd64 /usr/local/bin/healthcheck
RUN chmod +x /usr/local/bin/healthcheck

HEALTHCHECK CMD healthcheck
```

If doing a multiplatform build, you can also do:
```Dockerfile
FROM alpine
ARG TARGETARCH

ADD hhttps://github.com/jungsoft/docker-healthcheck/releases/download/0.1.0/healthcheck_alpine_$TARGETARCH /usr/local/bin/healthcheck
RUN chmod +x /usr/local/bin/healthcheck
```

Or if you prefer COPYing:

```Dockerfile
FROM alpine

# Assuming local file already has permission to execute
COPY local/path/to/healthcheck /usr/local/bin/healthcheck

HEALTHCHECK CMD healthcheck
```

## Building

To build it for alpine, you can use the go-alpine image, for example (alternatively, for ARM64, you can use `docker run --rm -it arm64v8/golang:1.17-alpine`):

```bash
docker run --rm -it --platform linux/amd64 amd64/golang:1.17-alpine
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
