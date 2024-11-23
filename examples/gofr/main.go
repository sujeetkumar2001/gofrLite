package main

import "gofr.dev/pkg/gofr"

func main(){
	a:=gofr.New()
	a.POST("/users",PostHanldler)

	a.Run()
}

func PostHanldler(ctx *gofr.Context)(interface{},error){

}

func GetByIdHanldler(ctx *gofr.Context)(interface{},error){
	
}