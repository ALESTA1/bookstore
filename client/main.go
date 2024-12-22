package main

import (
	"context"
	"log"
	"time"

	pb "bookstore/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	
	registerReq := &pb.RegisterRequest{
		Username: "testuser",
		Password: "testpassword",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = client.Register(ctx, registerReq)
	if err != nil {
		log.Fatalf("Failed to register: %v", err)
	}
	log.Println("User registered successfully")

	
	loginReq := &pb.LoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}
	loginResp, err := client.Login(ctx, loginReq)
	if err != nil {
		log.Fatalf("Failed to log in: %v", err)
	}
	token := loginResp.AccessToken

	log.Printf("Access token received: %s\n", token)

	
	authCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", token))

	
	createReq := &pb.CreateBookRequest{
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Genre:  "Fiction",
		Year:   "1951",
		Price:  12.99,
	}
	createResp, err := client.CreateBook(authCtx, createReq)
	if err != nil {
		log.Fatalf("Failed to create book: %v", err)
	}
	createdBook := createResp.Book
	log.Printf("Book created: %+v\n", createdBook)

	getReq := &pb.GetBookRequest{
		Id: createdBook.Id,
	}
	getResp, err := client.GetBook(authCtx, getReq)
	if err != nil {
		log.Fatalf("Failed to get book: %v", err)
	}
	log.Printf("Retrieved book: %+v\n", getResp.Book)

	
	updateReq := &pb.UpdateBookRequest{
		Book: &pb.Book{
			Id:     createdBook.Id,
			Title:  "The Catcher in the Rye (Updated)",
			Author: "J.D. Salinger (Updated)",
			Genre:  "Fiction (Updated)",
			Year:   "1951 (Updated)",
			Price:  15.99,
		},
	}
	updateResp, err := client.UpdateBook(authCtx, updateReq)
	if err != nil {
		log.Fatalf("Failed to update book: %v", err)
	}
	log.Printf("Updated book: %+v\n", updateResp.Book)

	
	listReq := &pb.ListBooksRequest{}
	listResp, err := client.ListBooks(authCtx, listReq)
	if err != nil {
		log.Fatalf("Failed to list books: %v", err)
	}
	log.Printf("List of books: %+v\n", listResp.Books)

	
	deleteReq := &pb.DeleteBookRequest{
		Id: createdBook.Id,
	}
	deleteResp, err := client.DeleteBook(authCtx, deleteReq)
	if err != nil {
		log.Fatalf("Failed to delete book: %v", err)
	}
	log.Printf("Deleted book success: %v\n", deleteResp.Success)

	
	listRespAfterDelete, err := client.ListBooks(authCtx, listReq)
	if err != nil {
		log.Fatalf("Failed to list books: %v", err)
	}
	log.Printf("List of books after deletion: %+v\n", listRespAfterDelete.Books)
}
