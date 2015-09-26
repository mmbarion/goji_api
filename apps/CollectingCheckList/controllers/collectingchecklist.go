package controllers

import (
	"net/http"
	"fmt"
)

func CCLHome(w http.ResponseWriter,req *http.Request) {
	fmt.Fprintf(w,"CCLHome")
}