package transport

import (
	"net/http"

	"lastlegends-gateway-service/middleware"

	"lastlegends-gateway-service/internal/service"

	"lastlegends-gateway-service/internal/endpoint"

	"lastlegends-gateway-service/helper"

	pb "github.com/ranggadablues/lastlegends-proto-library/user-proto/pb"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/gorilla/mux"
	"github.com/ranggadablues/gosok/auth"
	"github.com/ranggadablues/gosok/common"
)

func RegisterUserRoutes(mux *mux.Router, eps endpoint.UserEndpoints) {
	const usersId = "/users/{id}"

	publicRouter := mux.PathPrefix("/").Subrouter()

	// POST /users (create or login user)
	publicRouter.Methods("POST").Path("/users/{module}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userActionHandler(w, r, eps.Service)
	})

	// POST /refresh (refresh token)
	publicRouter.Methods("POST").Path("/refresh").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		refreshTokenHandler(w, r, eps.Service)
	})

	privateRouter := mux.PathPrefix("/").Subrouter()

	// GET /users (list users)
	privateRouter.Methods("GET").Path("/users").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getUsersHandler(w, r, eps.Service)
	})

	// GET /users/me (get user by JWT)
	privateRouter.Methods("GET").Path("/users/me").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getUserByMeHandler(w, r, eps.Service)
	})

	// GET /users/{id} (get user by id)
	privateRouter.Methods("GET").Path(usersId).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getUserByIdHandler(w, r, eps.Service)
	})

	// PUT /users/{id} (update user)
	privateRouter.Methods("PUT").Path(usersId).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		updateUserHandler(w, r, eps.Service)
	})

	// DELETE /users/{id} (delete user)
	privateRouter.Methods("DELETE").Path(usersId).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deleteUserHandler(w, r, eps.Service)
	})
	privateRouter.Use(middleware.AuthMiddleware)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request, userService service.IUserService) {
	req := &pb.ListUsersRequest{}
	resp, err := userService.GetUsers(r.Context(), req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	helper.WriteSuccessHeader(w, resp)
}

func getUserByIdHandler(w http.ResponseWriter, r *http.Request, userService service.IUserService) {
	vars := mux.Vars(r)
	req := &pb.GetUserByIdRequest{
		Id: vars["id"],
	}

	resp, err := userService.GetUserById(r.Context(), req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	helper.WriteSuccessHeader(w, resp)
}

func getUserByMeHandler(w http.ResponseWriter, r *http.Request, userService service.IUserService) {
	req := &pb.Empty{}
	ctx := auth.InjectToGRPCContext(r.Context())
	resp, err := userService.GetUserByMe(ctx, req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}
	helper.WriteSuccessHeader(w, resp)
}

func userActionHandler(w http.ResponseWriter, r *http.Request, userService service.IUserService) {
	vars := mux.Vars(r)
	var body pb.User
	if err := common.Payload(&body, r); err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	var response any
	var err error
	switch vars["module"] {
	case "register":
		req := &pb.CreateUserRequest{
			User: &body,
		}
		var resp *pb.CreateUserResponse
		resp, err = userService.CreateUser(r.Context(), req)
		response = resp.GetUser()
	case "login":
		req := &pb.LoginRequest{
			Email:    body.Email,
			Password: body.Password,
		}
		var resp *pb.LoginResponse
		resp, err = userService.LoginUser(r.Context(), req)
		response = bson.M{"token": resp.GetToken(), "refreshToken": resp.GetRefreshToken()}
	}

	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	helper.WriteSuccessHeader(w, response)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request, userService service.IUserService) {
	vars := mux.Vars(r)
	var payload bson.M
	if err := common.Payload(&payload, r); err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	body := pb.UserUpdate{}
	bsonBytes, err := bson.Marshal(payload)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}
	bson.Unmarshal(bsonBytes, &body)
	req := &pb.UpdateUserRequest{
		Id:   vars["id"],
		User: &body,
	}

	_, err = userService.UpdateUser(r.Context(), req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	helper.WriteSuccessHeader(w, nil)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request, userService service.IUserService) {
	vars := mux.Vars(r)
	req := &pb.DeleteUserRequest{
		Id: vars["id"],
	}

	resp, err := userService.DeleteUser(r.Context(), req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	helper.WriteSuccessHeader(w, resp)
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request, userService service.IUserService) {
	req := &pb.RefreshRequest{}
	resp, err := userService.RefreshToken(r.Context(), req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}
	helper.WriteSuccessHeader(w, resp)
}
