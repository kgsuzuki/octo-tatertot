FROM ioft/armhf-ubuntu

ENV GOPATH /go

RUN apt-get update && apt-get install -y golang nginx git && apt-get clean 

RUN go get github.com/stianeikeland/go-rpio

EXPOSE 80
