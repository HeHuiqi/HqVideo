package main

import (
	"github.com/julienschmidt/httprouter"
	"HqVideo/streamserver/limiter"
	"net/http"
	"fmt"
	"HqVideo/streamserver/handlers"
	"HqVideo/streamserver/response"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *limiter.ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router,cc int) http.Handler  {

	m := middleWareHandler{}
	m.r = r
	m.l = limiter.NewConnLimiter(cc)

	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {

	if !m.l.GetConn() {
		response.SendErrorResponse(w,http.StatusTooManyRequests,"Too many Request")
		return
	}
	m.r.ServeHTTP(w,r)
	defer m.l.ReleaseConn()
}
//注册处理器
func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()

	router.GET("/video/:vid-id",handlers.StreamHandler)
	router.POST("/upload/:vid-id",handlers.UploadHandler)
	router.GET("/",handlers.HqTestUploadHandler)

	return router
}

func main() {

	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r,2)
	fmt.Println("Video-stream:http://127.0.0.1:9000")
	http.ListenAndServe(":9000",mh)
}


