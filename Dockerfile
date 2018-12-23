FROM golang

MAINTAINER light4d
WORKDIR /GOPATH/src/github.com/light4d/object4d
ADD . /GOPATH/src/github.com/light4d/object4d
ENV GOPATH=/GOPATH
RUN cd /GOPATH/src/github.com/light4d/object4d && go get
EXPOSE 9001
VOLUME /GOPATH/src/github.com/light4d/object4d/bin/
CMD  ["/GOPATH/bin/object4d" "/GOPATH/src/github.com/light4d/object4d/bin/config.json"]
