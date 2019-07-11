package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "myLibrary/src/protoGetBookInfo"
	"time"
)

const (
	address = "dk.gribgp.com:50061"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGetBookInformationClient(conn)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second * 20)
	defer cancel()

	r1, err := c.GetBookInfoByID(ctx, &pb.IdRequest{Id: "27133480"})
	if err != nil {
		log.Fatalf("could not GetBookInfoByID 27133480: %v", err)
	}
	log.Printf("get book info json: %s", r1.InfoJson)
	log.Printf("get book info json ErrorMessage: %s", r1.ErrorMessage)

	//https://book.douban.com/subject/%s/
	r2, err := c.GetBookInfoByUrl(ctx, &pb.UrlRequest{Url: "https://book.douban.com/subject/27133480/"})
	if err != nil {
		log.Fatalf("could not GetBookInfoByID https://book.douban.com/subject/27133480/: %v", err)
	}
	log.Printf("get book info json: %s", r2.InfoJson)
	log.Printf("get book info json ErrorMessage: %s", r2.ErrorMessage)

}
