FROM golang:1.9-alpine AS build
COPY . /go/src/github.com/hulilabs/dom-api
WORKDIR /go/src/github.com/hulilabs/dom-api
RUN go build

FROM alpine:3.6
RUN apk add --update ca-certificates curl && rm -rf /var/cache/apk/*
HEALTHCHECK --interval=30s --retries=3 --timeout=5s \
    CMD curl --fail http://localhost:8000 || exit 1
COPY --from=build /go/src/github.com/hulilabs/dom-api/dom-api /dom-api
ENTRYPOINT ["/dom-api"]
EXPOSE 8000