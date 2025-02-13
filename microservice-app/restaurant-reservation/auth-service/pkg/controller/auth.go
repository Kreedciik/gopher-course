package controller

import (
	pb "auth/grpc_gen/auth"
	"auth/model"
	"auth/pkg/service"
	"context"
)

type AuthServer struct {
	user service.User
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer(user service.User) *AuthServer {
	return &AuthServer{user: user}
}

func (h *AuthServer) SignUp(ctx context.Context, req *pb.SignupReq) (*pb.SignupResp, error) {
	var newUser = model.CreateUserDTO{req.Username, req.Email, req.Password}
	if err := h.user.SignUp(ctx, newUser); err != nil {
		return nil, err
	}

	return &pb.SignupResp{Success: true}, nil
}

func (h *AuthServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	var credentials = model.SignInDTO{Email: req.Email, Password: req.Password}

	token, err := h.user.Login(ctx, credentials)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResp{Token: token}, nil
}

func (h *AuthServer) GetMe(ctx context.Context, req *pb.GetMeReq) (*pb.GetMeResp, error) {
	return nil, nil
}
