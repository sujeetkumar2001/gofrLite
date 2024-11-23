package main

import (
	"gofr.dev/pkg/gofr"
	"gofrHackathon/gofrLite"
)

func main() {
	a := gofr.New()
	a.POST("/users", gofrLite.Handler(createUser))
	a.GET("/users", gofrLite.Handler(listUsers))
	a.GET("/users/{id}", gofrLite.Handler(getUser))
	a.Run()
}

type User struct {
	ID     int
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	Age    int    `json:"age" validate:"gte=21"`
	Phone  string `json:"phone" validate:"required,e164"`
	Status string `json:"status" default:"active" validate:"oneof=active inactive"`
}

func createUser(ctx *gofr.Context, req *User) (*User, error) {
	// Handle create user logic
	return req, nil
}

type UserFilter struct {
	Age    int    `query:"age"`
	Status string `query:"status" validate:"oneof=active inactive"`
}

func listUsers(ctx *gofr.Context, req *UserFilter) ([]*User, error) {
	// Handle list users logic
	return []*User{{Age: req.Age, Status: req.Status}}, nil
}

type UserID struct {
	ID int `path:"id"`
}

func getUser(ctx *gofr.Context, req *UserID) (*User, error) {
	// Handle get user logic
	return &User{ID: req.ID}, nil
}
