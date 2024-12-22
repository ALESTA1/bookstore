package main

import (
	"bookstore/db"
	pb "bookstore/proto"
	"context"
	"fmt"
	"log"
)

func (s *helloServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {

	if err := ValidateDeleteBookRequest(req); err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	db.Mu.Lock()
	defer db.Mu.Unlock()

	bookID := req.Id
	_, exists := db.BookMap[bookID]
	if !exists {

		log.Printf("Book with id %d not found\n", bookID)
		return &pb.DeleteBookResponse{
			Success: false,
		}, nil
	}

	delete(db.BookMap, bookID)

	log.Printf("Book with id %d deleted successfully\n", bookID)

	return &pb.DeleteBookResponse{
		Success: true,
	}, nil
}
