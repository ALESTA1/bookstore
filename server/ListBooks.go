package main

import (
	"bookstore/db"
	pb "bookstore/proto"
	"context"
	"log"
)

func (s *helloServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {

	var books []*pb.Book
	for _, book := range db.BookMap {
		books = append(books, book)
	}

	log.Printf("Returning %d books\n", len(books))

	return &pb.ListBooksResponse{
		Books: books,
	}, nil
}
