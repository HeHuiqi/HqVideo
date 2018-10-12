package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"HqVideo/web/handlers"
	"fmt"
)

func RegisterHandler() *httprouter.Router  {

	router := httprouter.New()

	router.GET("/",handlers.HomeHandler)
	router.POST("/",handlers.HomeHandler)

	router.GET("/userhome",handlers.UserHomeHandler)
	router.POST("/userhome",handlers.UserHomeHandler)

	router.POST("/api",handlers.ApiHandler)

	router.POST("/upload/:vid-id",handlers.ProxyUploadHandler)


	router.ServeFiles("/static/*filepath",http.Dir("./web/templates/"))

	return router

}
func main() {

	r := RegisterHandler()

	fmt.Println("Video-web:http://127.0.0.1:8001")

	http.ListenAndServe(":8001",r)

}