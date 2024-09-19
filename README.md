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

To build it for alpine, you can use the go-alpine image, for example (alternatively, for ARM64, you can use `docker run --rm -it arm64v8/golang:1.17-alpine`) or directly from the command line with go installed:

```bash
go build healthcheck.go
```

The binary will be created inside the current directory with name `healthcheck`

Or alternatively for other operational systems or architectures, use respectively the envs `GOOS` and `GOARCH` (for better portability, make sure to use static linking with `CGO_ENABLED=0`):

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o healthcheck_linux_amd64 healthcheck.go

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o healthcheck_linux_arm64 healthcheck.go
```
