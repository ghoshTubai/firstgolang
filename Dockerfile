FROM golang:alpine
RUN mkdir -p $GOPATH/src/hello
COPY . $GOPATH/src/hello
WORKDIR $GOPATH/src/hello/
ENTRYPOINT go run project.go