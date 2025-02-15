package controller

import (
	pb "auth/grpc_gen/auth"
	"auth/model"
	"auth/pkg/service"
	"context"
)

type AuthServer struct {
	userService service.User
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer(userService service.User) *AuthServer {
	return &AuthServer{
		userService: userService,
	}
}

func (h *AuthServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	newUser := model.CreateUserDTO{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password}

	if err := h.userService.SignUp(newUser); err != nil {
		return nil, err
	}

	return &pb.SignUpResponse{Success: true}, nil
}

func (h *AuthServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	credentials := model.SignInDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	accessToken, err := h.userService.SignIn(credentials)
	if err != nil {
		return nil, err
	}

	return &pb.SignInResponse{Token: accessToken}, nil
}

func (h *AuthServer) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	userId := req.Id
	user, err := h.userService.GetUserProfile(userId)
	if err != nil {
		return nil, err
	}

	return &pb.GetProfileResponse{
		Id:    user.Id,
		Name:  user.Email,
		Email: user.Email,
	}, nil
}
