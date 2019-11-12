FROM golang:1.12

RUN printf "deb http://archive.debian.org/debian/ jessie main\ndeb-src http://archive.debian.org/debian/ jessie main\ndeb http://security.debian.org jessie/updates main\ndeb-src http://security.debian.org jessie/updates main" > /etc/apt/sources.list

RUN apt-get update && apt-get install -y netcat

RUN mkdir /go/src/github.com \
/go/src/github.com/alamyudi \
/go/src/github.com/alamyudi/echo-app \
/go/src/github.com/alamyudi/echo-app/echokit  \
/go/src/github.com/alamyudi/echo-app/mobilerestkit 

WORKDIR /go/src/github.com/alamyudi/echo-app/mobilerestkit

ADD glide.yaml /go/src/github.com/alamyudi/echo-app
ADD ./mobilerestkit/run-server.sh /go/src/github.com/alamyudi/echo-app/mobilerestkit

RUN curl https://glide.sh/get | sh

RUN rm -rf /go/src/github.com/alamyudi/echo-app/vendor
RUN cd /go/src/github.com/alamyudi/echo-app/ && glide install

RUN go get github.com/go-playground/locales
RUN go get github.com/oxequa/realize

RUN chmod +x /go/src/github.com/alamyudi/echo-app/mobilerestkit/run-server.sh

CMD ["sh", "run-server.sh", "mysql", "3306", "120"]
