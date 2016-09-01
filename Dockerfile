FROM golang:1.7
RUN export GOPATH=/go
RUN go get github.com/lib/pq
COPY . /go/src/github.com/stevenchung/microservice/
WORKDIR /go/src/github.com/stevenchung/microservice/
CMD [ "go", "run", "*.go" ]
