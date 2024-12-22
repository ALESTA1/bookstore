package main

import (
	pb "bookstore/proto"
	"bookstore/server/users"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *helloServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	err := ValidateLoginRequest(req)
	if err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	users.UserStoreLock.Lock()
	defer users.UserStoreLock.Unlock()

	username := req.Username
	password := req.Password

	_, exists := users.UserStore[username]

	if !exists {
		return nil, fmt.Errorf("user does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.UserStore[username].HashedPassword), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("invalid credentials %v", err)
	}

	accessToken, err := s.AuthSvc.IssueToken(ctx, username)

	if err != nil {
		return nil, fmt.Errorf("could not generate access token %v", err)
	}

	return &pb.LoginResponse{
		AccessToken: accessToken,
	}, nil
}
