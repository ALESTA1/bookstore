package main

import (
	pb "bookstore/proto"
	"errors"
)

func ValidateCreateBookRequest(req *pb.CreateBookRequest) error {
	if req.Title == "" {
		return errors.New("title cannot be empty")
	}
	if req.Author == "" {
		return errors.New("author cannot be empty")
	}
	if req.Genre == "" {
		return errors.New("genre cannot be empty")
	}
	if req.Year == "" {
		return errors.New("year must not be empty")
	}
	if req.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}

func ValidateGetBookRequest(req *pb.GetBookRequest) error {
	if req.Id <= 0 {
		return errors.New("id must be a positive integer")
	}
	return nil
}

func ValidateUpdateBookRequest(req *pb.UpdateBookRequest) error {
	if req.Book == nil {
		return errors.New("book cannot be nil")
	}
	if req.Book.Id <= 0 {
		return errors.New("id must be a positive integer")
	}
	if req.Book.Title == "" {
		return errors.New("title cannot be empty")
	}
	if req.Book.Author == "" {
		return errors.New("author cannot be empty")
	}
	if req.Book.Genre == "" {
		return errors.New("genre cannot be empty")
	}
	if req.Book.Year == "" {
		return errors.New("year cannot be empty")
	}
	if req.Book.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}

func ValidateDeleteBookRequest(req *pb.DeleteBookRequest) error {
	if req.Id <= 0 {
		return errors.New("id must be a positive integer")
	}
	return nil
}

func ValidateRegisterRequest(req *pb.RegisterRequest) error {
	if req.Username == "" {
		return errors.New("cannot be empty")
	}
	if req.Password == "" {
		return errors.New("cannot be empty")
	}
	return nil
}

func ValidateLoginRequest(req *pb.LoginRequest) error {
	if req.Username == "" {
		return errors.New("cannot be empty")
	}
	if req.Password == "" {
		return errors.New("cannot be empty")
	}
	return nil
}
