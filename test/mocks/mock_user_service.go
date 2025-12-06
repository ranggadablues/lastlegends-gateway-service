package mocks

import (
	"context"

	pb "github.com/ranggadablues/lastlegends-proto-library/user-proto/pb"
)

type MockUserService struct {
	GetUsersFn     func(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error)
	GetUserByIdFn  func(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error)
	GetUserByMeFn  func(ctx context.Context, req *pb.Empty) (*pb.AuthResponse, error)
	CreateUserFn   func(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	UpdateUserFn   func(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUserFn   func(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	LoginUserFn    func(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	RefreshTokenFn func(ctx context.Context, req *pb.RefreshRequest) (*pb.LoginResponse, error)
}

func (m *MockUserService) GetUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return m.GetUsersFn(ctx, req)
}
func (m *MockUserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	return m.GetUserByIdFn(ctx, req)
}
func (m *MockUserService) GetUserByMe(ctx context.Context, req *pb.Empty) (*pb.AuthResponse, error) {
	return m.GetUserByMeFn(ctx, req)
}
func (m *MockUserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return m.CreateUserFn(ctx, req)
}
func (m *MockUserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return m.UpdateUserFn(ctx, req)
}
func (m *MockUserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return m.DeleteUserFn(ctx, req)
}
func (m *MockUserService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return m.LoginUserFn(ctx, req)
}
func (m *MockUserService) RefreshToken(ctx context.Context, req *pb.RefreshRequest) (*pb.LoginResponse, error) {
	return m.RefreshTokenFn(ctx, req)
}
