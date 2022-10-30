FROM golang:1.8 as goimage

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /root/

COPY . app/
WORKDIR app/

ENV PORT 9090

RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./.bin/app ./cmd/api/main.go

FROM alpine:3.6 as baseimagealp
RUN apk add --no-cache bash
ENV WORK_DIR=/docker/bin
WORKDIR $WORK_DIR
COPY --from=goimage /go/src/github.com/aditmayapada/tryout/bin ./
ENTRYPOINT /docker/bin/main
EXPOSE 9090