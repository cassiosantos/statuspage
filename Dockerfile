FROM golang:alpine3.7 as builder

WORKDIR /go/src/github.com/involvestecnologia/statuspage

RUN apk add git --no-cache

COPY . .

RUN go get .

RUN GOOS=linux go build -o statuspage

FROM alpine:3.7

RUN addgroup -S statuspage && adduser -S -g statuspage statuspage

USER statuspage

COPY --from=builder /go/src/github.com/involvestecnologia/statuspage/statuspage /usr/bin/statuspage

EXPOSE 8080

CMD ["statuspage"]