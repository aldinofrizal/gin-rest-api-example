package services

import (
	"os"

	pb "github.com/aldinofrizal/gin-rest-api-example/services/news"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewsService() (pb.NewsServiceClient, error) {
	conn, err := grpc.Dial(os.Getenv("NEWS_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := pb.NewNewsServiceClient(conn)
	return c, nil
}
