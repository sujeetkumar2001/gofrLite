# gofrLite

gofrLite lets you write simpler [gofr](https://gofr.dev/) handlers.

**key features**
- Binds request body, query params and path variable into single request object.
- Reduces boilerplate code.

## Getting started
Add gofrLite to your project
````bash
go get github.com/sujeetkumar2001/gofrLite
````
Register paths with gofrLite Handler
````go
package main

import (
	"github.com/sujeetkumar2001/gofrLite"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.POST("/users", gofrLite.Handler(postUser))
	app.GET("/users/{id}", gofrLite.Handler(getUser))
	app.DELETE("/users/{id}", gofrLite.Handler(deleteUser))

	app.Run()
}
````
Define handler input objects
````go
package main

type PostUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=21"`
}

type GetUserRequest struct {
	ID    int    `path:"id"`
	Name  string `query:"name"`
	Email string `query:"email"`
	Age   int    `query:"age" default:"21"`
}

type DeleteUserRequest struct {
	ID int `path:"id"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
````
Implement gofrLite compatible handlers
````go
package main

import (
	"gofr.dev/pkg/gofr"
)

func postUser(ctx *gofr.Context, user *PostUserRequest) (*User, error) {
	// handle create user logic  
}

func getUser(ctx *gofr.Context, req *GetUserRequest) (*User, error) {
	// handle get user logic
}

func deleteUser(ctx *gofr.Context, req *DeleteUserRequest) error {
	// handle delete user logic
}
````