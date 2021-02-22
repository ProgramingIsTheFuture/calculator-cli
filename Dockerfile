FROM golang:1.15

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLE=1

RUN apt-get update && \
	go get github.com/alexflint/go-arg && \
	go get github.com/stretchr/testify 

cmd ["tail", "-f", "/dev/null"]
