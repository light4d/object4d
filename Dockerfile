FROM golang

MAINTAINER light4d
ADD . /GOPATH/src/github.com/light4d/object4d
ENV GOPATH=/GOPATH
RUN cd /GOPATH/src/github.com/light4d/object4d && go get
EXPOSE 9001
VOLUME /root/home/config.json
CMD  ["/GOPATH/bin/object4d" ]
