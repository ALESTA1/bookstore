package main

import (
	"bookstore/db"
	pb "bookstore/proto"
	"context"
	"fmt"
	"log"
)

func (s *helloServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {

	if err := ValidateCreateBookRequest(req); err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	db.Mu.Lock()
	defer db.Mu.Unlock()

	newID := int32(len(db.BookMap) + 1)

	newBook := &pb.Book{
		Id:     newID,
		Title:  req.Title,
		Author: req.Author,
		Genre:  req.Genre,
		Year:   req.Year,
		Price:  req.Price,
	}

	db.BookMap[newID] = newBook

	log.Printf("Book created: %+v\n", newBook)

	return &pb.CreateBookResponse{
		Book: newBook,
	}, nil
}
