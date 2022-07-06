package router

import (
	"context"
	"fmt"
	"github.com/SherzodAbdullajonov/service1/application"
	"github.com/SherzodAbdullajonov/service1/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	app *application.App
	pb.UnimplementedDataServiceServer
}

func NewGrpcServer(app *application.App) GrpcServer {
	return GrpcServer{
		app: app,
	}
}

func (s GrpcServer) DownloadPosts(ctx context.Context, _ *pb.DownloadPostsRequest) (*emptypb.Empty, error) {
	
	


	go func() {
		err := s.app.DownloadPosts(context.Background())
		if err != nil {
			fmt.Print("setting operations status to unsuccessfull")
		}

		if updateErr := s.app.SetDownloadStatus(context.Background(), err == nil, err); updateErr != nil {
			fmt.Errorf("failed to set download status: %v", updateErr)
		}
	}()

	return &emptypb.Empty{}, nil
}


