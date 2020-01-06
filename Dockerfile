FROM golang:latest

WORKDIR $GOPATH/src/FirstGin
COPY . $GOPATH/src/FirstGin
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./FirstGin"]