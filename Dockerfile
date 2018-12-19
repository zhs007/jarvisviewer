FROM golang:1.10 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR $GOPATH/src/github.com/zhs007/jarvisviewer

COPY ./Gopkg.* $GOPATH/src/github.com/zhs007/jarvisviewer/

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure -vendor-only -v

COPY . $GOPATH/src/github.com/zhs007/jarvisviewer

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jarvisviewer . \
    && mkdir /home/jarvisviewer \
    && mkdir /home/jarvisviewer/cfg \
    && mkdir /home/jarvisviewer/dat \
    && mkdir /home/jarvisviewer/logs \
    && cp ./jarvisviewer /home/jarvisviewer/ \
    && cp ./test /home/jarvisviewer/test \
    && cp ./cfg/config.yaml.default /home/jarvisviewer/cfg/config.yaml

FROM alpine
RUN apk upgrade && apk add --no-cache ca-certificates
WORKDIR /home/jarvisviewer
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /home/jarvisviewer /home/jarvisviewer
CMD ["./jarvisviewer"]
