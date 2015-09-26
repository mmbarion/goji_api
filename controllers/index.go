package controllers

import (
	"fmt"
	"net/http"
//	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)


func Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w,"Hello!")
}

func About(c web.C,w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w,"Hello, %s!",c.URLParams["name"])
}