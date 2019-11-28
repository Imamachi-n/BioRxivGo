FROM golang:1.13.4 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/Imamachi-n/BioRxivGo/server
COPY . .
RUN make docker

# runtime image
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/Imamachi-n/BioRxivGo/server/bin/goServer /goServer
EXPOSE 9000
ENTRYPOINT ["/goServer"]
