package service

import (
	"context"

	pb "github.com/ranggadablues/lastlegends-proto-library/user-proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IUserService interface {
	GetUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error)
	GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error)
	GetUserByMe(ctx context.Context, req *pb.Empty) (*pb.AuthResponse, error)
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	RefreshToken(ctx context.Context, req *pb.RefreshRequest) (*pb.LoginResponse, error)
}

type userService struct {
	client pb.UserServiceClient
}

func NewUserService(addr string) IUserService {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	return &userService{client: pb.NewUserServiceClient(conn)}
}

func (s *userService) GetUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return s.client.ListUsers(ctx, req)
}

func (s *userService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	return s.client.GetUserById(ctx, req)
}

func (s *userService) GetUserByMe(ctx context.Context, req *pb.Empty) (*pb.AuthResponse, error) {
	return s.client.GetUser(ctx, req)
}

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.client.CreateUser(ctx, req)
}

func (s *userService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.client.LoginUser(ctx, req)
}

func (s *userService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return s.client.UpdateUser(ctx, req)
}

func (s *userService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.client.DeleteUser(ctx, req)
}

func (s *userService) RefreshToken(ctx context.Context, req *pb.RefreshRequest) (*pb.LoginResponse, error) {
	return s.client.RefreshToken(ctx, req)
}
