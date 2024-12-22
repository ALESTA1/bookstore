package main

import (
	"bookstore/db"
	pb "bookstore/proto"
	"context"
	"fmt"
	"log"
)

func (s *helloServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {

	if err := ValidateUpdateBookRequest(req); err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	
	db.Mu.Lock()
	defer db.Mu.Unlock()
	
	bookID := req.Book.Id

	existingBook, exists := db.BookMap[bookID]
	if !exists {
		return nil, fmt.Errorf("book with id %d not found", bookID)
	}

	existingBook.Title = req.Book.Title
	existingBook.Author = req.Book.Author
	existingBook.Genre = req.Book.Genre
	existingBook.Year = req.Book.Year
	existingBook.Price = req.Book.Price

	log.Printf("Updated book: %+v\n", existingBook)

	return &pb.UpdateBookResponse{
		Book: existingBook,
	}, nil
}
