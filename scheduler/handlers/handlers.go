package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"HqVideo/scheduler/dbops"
	"HqVideo/scheduler/response"
)

func VideoDeleteRecordHandler(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {


	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		response.SendResponse(w,400,"video id should not be empty")
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		response.SendResponse(w,500,"Internal sever error")
		return
	}

	response.SendResponse(w,200,"delete file successful")



}