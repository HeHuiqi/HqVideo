package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"HqVideo/api/handlers"
	"HqVideo/api/session"
	"fmt"
	"log"
	"HqVideo/api/auth"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler  {
	m := middleWareHandler{}
	m.r = r
	return m
}
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {

	//在这里检测session的合法是否过期等
	auth.ValidateUserSession(r)
	m.r.ServeHTTP(w,r)
}
//注册处理器
func RegisterHandlers() *httprouter.Router  {


	router := httprouter.New()

	router.POST("/user",handlers.CreateUser)

	router.POST("/user/:username",handlers.Login)

	router.GET("/user/:username",handlers.HqGetUserInfo)

	router.POST("/user/:username/videos",handlers.AddNewVideo)

	router.GET("/user/:username/videos",handlers.ListAllVideo)

	router.DELETE("/user/:username/videos/:vid-id",handlers.DeleteVideo)

	router.POST("/videos/:vid-id/comments",handlers.PostComment)
	router.GET("/videos/:vid-id/comments",handlers.ShowComments)



	router.GET("/",handlers.HomePage)

	router.ServeFiles("/static/*filepath", http.Dir("./api/templates/"))

	return router

}
func Prepare()  {
	session.LoadSessionFromDB()
}
func main() {

	fmt.Println("Video:http://127.0.0.1:8000")

	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	log.Fatalln(http.ListenAndServe(":8000",mh))



}