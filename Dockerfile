ARG GOVERSION=1.17
FROM golang:${GOVERSION} as build-env

ARG GOARCH
ENV GOARCH=${GOARCH}

WORKDIR /go/src/gh-contrib
COPY . /go/src/gh-contrib

RUN go mod download
RUN go mod verify
RUN make build

# Copy into base image
FROM alpine:3.7
WORKDIR /opt/app
COPY --from=build-env ./go/src/gh-contrib/gh-contrib /opt/app

RUN adduser -s /bin/false -h /home/gh-contrib -D gh-contrib

USER gh-contrib

ENTRYPOINT ["/opt/app/gh-contrib"]

CMD ["help"]
