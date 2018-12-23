FROM golang

MAINTAINER light4d
ADD . /GOPATH/src/github.com/light4d/object4d
ADD bin/config.json /GOPATH/bin/config.json
ENV GOPATH=/GOPATH
RUN cd /GOPATH/src/github.com/light4d/object4d && go get
EXPOSE 9001
VOLUME /GOPATH/bin/config.json
CMD  ["/GOPATH/bin/object4d" "/GOPATH/bin/config.json"]
