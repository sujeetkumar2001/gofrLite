package main

import (
	"regexp"
	"strconv"

	"gofr.dev/pkg/gofr"
	gofrErr "gofr.dev/pkg/gofr/http"
)

type user struct{
	ID     int
	Name string `json:"name"`
    Email string `json:"email"`
	Age   int     `json:"age"`
	Phone string `json:"phone"`
	Status string `json:"status"`
}

type status string

const (
	ACTIVE status ="ACTIVE"
	INACTIVE status="INACTIVE"
)

func main(){
	a:=gofr.New()

	a.POST("/users",PostHanldler)
	a.GET("/users",GetHandler)
	a.GET("/users/{id}",GetByIDHanldler)

	a.Run()
}

func PostHanldler(ctx *gofr.Context)(interface{},error){
    user:=&user{}

	err:=ctx.Bind(user)
	if err!=nil{
		ctx.Logger.Error("Error biding request body")
		return nil ,err
	}

	if user.Name==""{
		return nil, gofrErr.ErrorMissingParam{Params:[]string{"name"}}
	}

	if user.Email==""{
		return nil, gofrErr.ErrorMissingParam{Params:[]string{"email"}}
	}

	if !isValidEmail(user.Email){
		return nil, gofrErr.ErrorInvalidParam{Params:[]string{"email"}}
	}

	if user.Age<=21{
		return nil, gofrErr.ErrorInvalidParam{Params:[]string{"age"}}
	}

	if user.Phone==""{
		return nil, gofrErr.ErrorMissingParam{Params:[]string{"phone"}}
	}

	if !isValidPhone(user.Phone){
		return nil, gofrErr.ErrorInvalidParam{Params:[]string{"phone"}}
	}

	if !isValidStatus(user.Status){
		return nil, gofrErr.ErrorInvalidParam{Params:[]string{"age"}}
	}


	return user,nil

}

func GetByIDHanldler(ctx *gofr.Context)(interface{},error){
	id:=ctx.PathParam("id")

	if id==""{
		return nil, gofrErr.ErrorMissingParam{Params:[]string{"id"}}
	}
   i,_:=strconv.Atoi(id)

	return &user{
		ID: i,
      Name: "Maji",
	  Email: "maji@test.com",
	  Age: 25,
	  Phone: "+916923482539",
	  Status: "ACTIVE",
	},nil
}

func GetHandler(ctx *gofr.Context)(interface{},error){
	age:=ctx.Param("age")
	if age==""{
		return nil, gofrErr.ErrorMissingParam{Params:[]string{"age"}}
	}

	status:=ctx.Param("status")
	if status==""{
		return nil, gofrErr.ErrorMissingParam{Params:[]string{"status"}}
	}

	if !isValidStatus(status){
		return nil, gofrErr.ErrorInvalidParam{Params:[]string{"status"}}
	}

	a,_:=strconv.Atoi(age)
	return &user{
		ID: 5,
		Name: "Maji",
		Email: "maji@test.com",
		Age:a ,
		Phone: "+916923482539",
		Status: status,
	  },nil
}

func isValidStatus(s string) bool {
	switch status(s) {
	case ACTIVE, INACTIVE:
		return true
	default:
		return false
	}
}

func isValidEmail(email string)bool{
	// Regular expression for validating email addresses
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(regex)

	// Match the email against the regular expression
	return re.MatchString(email)
}

func isValidPhone(phone string) bool {
	// Regular expression for validating Indian phone numbers
	regex := `^(\+91[-\s]?)?[6-9][0-9]{9}$`

	// Compile the regular expression
	re := regexp.MustCompile(regex)

	// Match the phone number against the regular expression
	return re.MatchString(phone)
}