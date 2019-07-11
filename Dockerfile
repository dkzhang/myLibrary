FROM dkzhang007/my-grpc-golang:my1.0

RUN go get github.com/PuerkitoBio/goquery

WORKDIR /go/src/myLibrary

COPY . /go/src/myLibrary

#RUN protoc -I queryDouBanID/ queryDouBanID/the.proto --go_out=plugins=grpc:queryDouBanID/

RUN go build ./src/main.go

CMD ["./main"]