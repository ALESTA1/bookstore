package main

import (
	"context"
	"log"
	"testing"
	"time"

	pb "bookstore/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"net"
)


func startTestServer() *grpc.Server {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	reflection.Register(server)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return server
}

func TestClientServerInteraction(t *testing.T) {
	
	server := startTestServer()
	defer server.Stop() 
	
	time.Sleep(1 * time.Second)

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to connect to server: %v", err)
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
	assert.NoError(t, err, "Failed to register user")

	
	loginReq := &pb.LoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}
	loginResp, err := client.Login(ctx, loginReq)
	assert.NoError(t, err, "Failed to login")
	assert.NotEmpty(t, loginResp.AccessToken, "Access token should not be empty")
	token := loginResp.AccessToken

	
	authCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", token))
	createReq := &pb.CreateBookRequest{
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Genre:  "Fiction",
		Year:   "1951",
		Price:  12.99,
	}
	createResp, err := client.CreateBook(authCtx, createReq)
	assert.NoError(t, err, "Failed to create book")
	assert.Equal(t, "The Catcher in the Rye", createResp.Book.Title, "Book title should match")

	getReq := &pb.GetBookRequest{
		Id: createResp.Book.Id,
	}
	getResp, err := client.GetBook(authCtx, getReq)
	assert.NoError(t, err, "Failed to get book")
	assert.Equal(t, "The Catcher in the Rye", getResp.Book.Title, "Retrieved book title should match")

	
	updateReq := &pb.UpdateBookRequest{
		Book: &pb.Book{
			Id:     createResp.Book.Id,
			Title:  "The Catcher in the Rye (Updated)",
			Author: "J.D. Salinger (Updated)",
			Genre:  "Fiction (Updated)",
			Year:   "1951 (Updated)",
			Price:  15.99,
		},
	}
	updateResp, err := client.UpdateBook(authCtx, updateReq)
	assert.NoError(t, err, "Failed to update book")
	assert.Equal(t, "The Catcher in the Rye (Updated)", updateResp.Book.Title, "Updated book title should match")


	listReq := &pb.ListBooksRequest{}
	listResp, err := client.ListBooks(authCtx, listReq)
	assert.NoError(t, err, "Failed to list books")
	assert.Len(t, listResp.Books, 1, "There should be 1 book in the list")

	deleteReq := &pb.DeleteBookRequest{
		Id: createResp.Book.Id,
	}
	deleteResp, err := client.DeleteBook(authCtx, deleteReq)
	assert.NoError(t, err, "Failed to delete book")
	assert.True(t, deleteResp.Success, "Book deletion should be successful")

	listRespAfterDelete, err := client.ListBooks(authCtx, listReq)
	assert.NoError(t, err, "Failed to list books after deletion")
	assert.Len(t, listRespAfterDelete.Books, 0, "There should be no books after deletion")
}
