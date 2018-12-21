
FROM golang
WORKDIR $GOPATH/src/golang-demo
ADD . $GOPATH/src/golang-demo
RUN go build .
CMD ./golang-demo
