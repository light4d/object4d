FROM golang

MAINTAINER light4d
WORKDIR /GOPATH/src/github.com/light4d/object4d
ADD . /GOPATH/src/github.com/light4d/object4d
RUN cd /GOPATH/src/github.com/light4d/object4d && go build .
EXPOSE 9001
CMD  ["./object4d" ]
