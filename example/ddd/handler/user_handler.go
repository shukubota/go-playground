package handler

import (
	"fmt"
	"net/http"
)

//// usecaseとか注入する
//type userHandler struct{}
//
//func NewUserHandler() *userHandler {
//	return &userHandler{}
//}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get_user")
	fmt.Println(w)
	fmt.Println(r)
}
