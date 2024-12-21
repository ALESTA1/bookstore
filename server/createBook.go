package main

import (
	"bookstore/db"
	pb "bookstore/proto"
	"context"
	"fmt"
	"log"
)

func (s *helloServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {

	db.Mu.Lock()
	defer db.Mu.Unlock()

	newID := int32(len(db.BookMap) + 1)

	newBook := &pb.Book{
		Id:     newID,
		Title:  req.Title,
		Author: req.Author,
		Genre:  req.Genre,
		Year:   fmt.Sprintf("%d", req.Year),
		Price:  req.Price,
	}

	db.BookMap[newID] = newBook

	log.Printf("Book created: %+v\n", newBook)

	return &pb.CreateBookResponse{
		Book: newBook,
	}, nil
}
