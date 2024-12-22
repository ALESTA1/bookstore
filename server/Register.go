package main

import (
	pb "bookstore/proto"
	"bookstore/server/users"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *helloServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	err := ValidateRegisterRequest(req)
	if err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	users.UserStoreLock.Lock()
	defer users.UserStoreLock.Unlock()

	username := req.Username
	password := req.Password

	_, exists := users.UserStore[username]
	if exists {
		return &pb.RegisterResponse{
			Success: false,
		}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return &pb.RegisterResponse{
			Success: false,
		}, nil
	}

	users.UserStore[username] = &users.User{
		Username:       username,
		HashedPassword: string(hashedPassword),
	}

	return &pb.RegisterResponse{
		Success: true,
	}, nil
}
