FROM dkzhang007/my-grpc-golang:my1.0

RUN go get github.com/PuerkitoBio/goquery

WORKDIR /go/src/myLibrary

COPY . /go/src/myLibrary

RUN protoc -I src/protoGetBookInfo/ src/protoGetBookInfo/GetBookInfo.proto --go_out=plugins=grpc:src/protoGetBookInfo/

RUN go build ./src/main.go

CMD ["./main"]

# docker run --name bookinfo-spider -p 50061:50061 --mount source=book-cover,target=/BookCover dkzhang007/book-library-bookinfo-spider:latest