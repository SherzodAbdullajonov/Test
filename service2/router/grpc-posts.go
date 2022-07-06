package router

import (
	"context"
	"fmt"

	"github.com/SherzodAbdullajonov/service2/application"
	"github.com/SherzodAbdullajonov/service2/pb"
	"github.com/SherzodAbdullajonov/service2/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	app application.App
	pb.UnimplementedPostServiceServer
}

func NewGrpcServer(app application.App) GrpcServer {
	return GrpcServer{
		app: app,
	}
}

func (s GrpcServer) GetPost(ctx context.Context, request *pb.GetPostRequest) (*pb.Post, error) {
	p, err := s.app.Queries.GetPost.Handle(ctx, int(request.PostId))
	if err != nil {
		return &pb.Post{}, err
	}

	return toProtoPost(p), nil
}

func (s GrpcServer) ListPosts(ctx context.Context, request *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := s.app.Queries.ListPosts.Handle(ctx, int(request.Page), int(request.Limit))
	if err != nil {
		return &pb.ListPostsResponse{}, err
	}

	protoPosts := make([]*pb.Post, 0, len(posts))
	for _, p := range posts {
		protoPosts = append(protoPosts, toProtoPost(p))
	}

	return &pb.ListPostsResponse{
		Posts: protoPosts,
	}, nil
}

func (s GrpcServer) UpdatePost(ctx context.Context, request *pb.Post) (*emptypb.Empty, error) {
	p, err := repository.New(int(request.Id), int(request.UserId), request.Title, request.Body)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	err = s.app.Commands.UpdatePost.Handle(ctx, p)
	fmt.Printf("Update err=%v", err)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s GrpcServer) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*emptypb.Empty, error) {
	err := s.app.Commands.DeletePost.Handle(ctx, int(request.PostId))

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func toProtoPost(p repository.Post) *pb.Post {
	return &pb.Post{
		Id:     int64(p.ID()),
		UserId: int64(p.UserID()),
		Title:  p.Title(),
		Body:   p.Body(),
	}
}
