FROM dkzhang007/my-grpc-golang:my1.0

RUN go get github.com/PuerkitoBio/goquery

WORKDIR /go/src/myLibrary

COPY . /go/src/myLibrary

RUN protoc -I protoGetBookInfo/ protoGetBookInfo/GetBookInfo.proto --go_out=plugins=grpc:protoGetBookInfo/

RUN go build ./src/main.go

CMD ["./main"]