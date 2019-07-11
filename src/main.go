package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"myLibrary/src/bookInfo"
	pb "myLibrary/src/protoGetBookInfo"
	"net"
)

//func ExampleScrape() {
//
//	theBookInfo := bookInfo.BookInfo{}
//	//theBookInfo.parseFromHtmlDouBan("27015617")
//	//theBookInfo.parseFromHtmlDouBan("27665114")
//	theBookInfo.ParseFromHtmlDouBanID("27133480", true)
//
//	log.Printf("%v", theBookInfo.TheBookBasicInfo)
//	log.Printf("%v", theBookInfo.TheDouBanRating)
//
//	log.Printf("ContentIntroduce: %s", theBookInfo.TheBookIntroduce.ContentIntroduce)
//	log.Printf("AuthorIntroduce: %s", theBookInfo.TheBookIntroduce.AuthorIntroduce)
//	log.Printf("Tags: %s", theBookInfo.TheBookTagAndRec.Tags)
//}

const (
	port = ":50061"
)

type server struct{}

func (s *server) GetBookInfoByID(ctx context.Context, in *pb.IdRequest) (*pb.BookInfoJsonReply, error) {
	log.Printf("Received: %v", in.Id)

	theBookInfo := bookInfo.BookInfo{}
	theBookInfo.ParseFromHtmlDouBanID(in.Id, true)

	infoJson, err := json.MarshalIndent(theBookInfo, "", "    ")
	if err != nil {
		log.Printf("json.MarshalIndent error: %v", err)

		return &pb.BookInfoJsonReply{
			InfoJson:     "",
			ErrorMessage: fmt.Sprintf("json.MarshalIndent error: %v", err),
		}, nil
	}
	return &pb.BookInfoJsonReply{
		InfoJson:     string(infoJson),
		ErrorMessage: "",
	}, nil
}

func (s *server) GetBookInfoByUrl(ctx context.Context, in *pb.UrlRequest) (*pb.BookInfoJsonReply, error) {
	log.Printf("Received: %v", in.Url)

	theBookInfo := bookInfo.BookInfo{}
	theBookInfo.ParseFromHtmlDouBanURL(in.Url, true)

	infoJson, err := json.MarshalIndent(theBookInfo, "", "    ")
	if err != nil {
		log.Printf("json.MarshalIndent error: %v", err)

		return &pb.BookInfoJsonReply{
			InfoJson:     "",
			ErrorMessage: fmt.Sprintf("json.MarshalIndent error: %v", err),
		}, nil
	}
	return &pb.BookInfoJsonReply{
		InfoJson:     string(infoJson),
		ErrorMessage: "",
	}, nil
}

func main() {
	//ExampleScrape()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetBookInformationServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
