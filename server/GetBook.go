package main

import (
	"bookstore/db"
	pb "bookstore/proto"
	"context"
	"fmt"
)

func (s *helloServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {

	if err := ValidateGetBookRequest(req); err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	book, exists := db.BookMap[req.Id]
	if !exists {
		return nil, fmt.Errorf("book with id %d not found", req.Id)
	}

	return &pb.GetBookResponse{
		Book: book,
	}, nil
}
