FROM golang:1.19

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN mkdir -p /opt/unix/workspace
ENV SOCKET_PATH=/opt/unix

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/client/main.go

USER nobody

CMD ["app"]
