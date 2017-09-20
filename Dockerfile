FROM golang:1.9-alpine AS build
COPY . /go/src/github.com/hulilabs/dom-api
WORKDIR /go/src/github.com/hulilabs/dom-api
RUN go build

FROM alpine:3.6
COPY --from=build /go/src/github.com/hulilabs/dom-api/dom-api /dom-api
ENTRYPOINT ["/dom-api"]
EXPOSE 8000