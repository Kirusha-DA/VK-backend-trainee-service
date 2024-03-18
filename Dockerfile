FROM golang:1.21.3-alpine
WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o /bin/server ./cmd/filmoteka
ENTRYPOINT [ "/bin/server" ]