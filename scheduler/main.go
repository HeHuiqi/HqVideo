package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"HqVideo/scheduler/taskrunner"
	"HqVideo/scheduler/handlers"
)

func RegisterHandlers() *httprouter.Router  {

	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id",handlers.VideoDeleteRecordHandler)
	router.GET("/",handlers.VideoSchedulerIndex)
	return router
}
func main() {
	fmt.Println("Video-scheduler:http://127.0.0.1:9001")

	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001",r)

}
