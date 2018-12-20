FROM golang:1.11.3-alpine3.8

RUN apk add --no-cache git

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

RUN go install -v ./...

CMD ["app"]