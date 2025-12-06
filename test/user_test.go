package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"lastlegends-gateway-service/internal/endpoint"
	"lastlegends-gateway-service/internal/transport"
	"lastlegends-gateway-service/test/mocks"

	"github.com/gorilla/mux"
	pb "github.com/ranggadablues/lastlegends-proto-library/user-proto/pb"
)

func TestGetUserSuccess(t *testing.T) {
	const firstname = "Test User"
	mockUserService := &mocks.MockUserService{
		GetUsersFn: func(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
			return &pb.ListUsersResponse{Users: []*pb.User{{Id: "123", Firstname: firstname}}}, nil
		},
		GetUserByIdFn: func(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
			return &pb.GetUserByIdResponse{User: &pb.User{Id: "123", Firstname: firstname}}, nil
		},
		CreateUserFn: func(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
			return &pb.CreateUserResponse{User: &pb.User{Id: "123", Firstname: firstname}}, nil
		},
		UpdateUserFn: func(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
			return &pb.UpdateUserResponse{Success: true}, nil
		},
		DeleteUserFn: func(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
			return &pb.DeleteUserResponse{Success: true}, nil
		},
	}

	router := mux.NewRouter()
	eps := endpoint.UserEndpoints{Service: mockUserService}
	transport.RegisterUserRoutes(router, eps)

	req := httptest.NewRequest("GET", "/api/user/123", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	if response["name"] != "Test User" {
		t.Fatalf("unexpected user name: %v", response["name"])
	}
}
