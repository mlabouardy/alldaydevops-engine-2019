FROM golang:1.10.1
WORKDIR /go/src/github.com/mlabouardy/imdb-engine
COPY main.go .
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
LABEL Maintainer mlabouardy
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/mlabouardy/imdb-engine/app .
CMD ["./app"] 