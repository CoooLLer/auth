package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/coooller/auth/pkg/user_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Get(_ context.Context, in *desc.GetUserIn) (*desc.GetUserOut, error) {
	log.Printf("Received request for get user: %d", in.Id)

	return &desc.GetUserOut{
		Id:        in.Id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Create(_ context.Context, in *desc.CreateUserIn) (*desc.CreateUserOut, error) {
	log.Printf("Received request to create user")
	log.Println(in)
	return &desc.CreateUserOut{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Update(_ context.Context, in *desc.UpdateUserIn) (*emptypb.Empty, error) {
	log.Printf("Received request to update user")
	log.Println(in)
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(_ context.Context, in *desc.DeleteUserIn) (*emptypb.Empty, error) {
	log.Printf("Received request to delete user: %d", in.Id)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
