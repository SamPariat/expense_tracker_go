ARG GO_VERSION=1.21.6

FROM golang:${GO_VERSION}

WORKDIR /server

COPY go.mod go.sum ./

COPY *.go ./

RUN go mod download

RUN cd cmd/server && go build -o ../../tmp/main.exe

EXPOSE 9623

ENTRYPOINT /tmp/main
