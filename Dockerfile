FROM golang:1.15

WORKDIR /go/src/among-go
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

CMD ["among-go"]