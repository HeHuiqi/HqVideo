package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"HqVideo/streamserver/defs"
	"os"
	"HqVideo/streamserver/response"
	"time"
	"log"
	"io/ioutil"
	"io"
	"html/template"
)

func StreamHandler(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {

	vid := p.ByName("vid-id")
	vl := defs.VIDEO_DIR +vid

	video,err := os.Open(vl)

	if err != nil {
		response.SendErrorResponse(w,http.StatusInternalServerError,"Internal error")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")

	//把文件返回到client
	http.ServeContent(w,r,"",time.Now(),video)
	defer video.Close()

}
func UploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

	r.Body = http.MaxBytesReader(w,r.Body,defs.MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(defs.MAX_UPLOAD_SIZE); err != nil {
		response.SendErrorResponse(w,http.StatusBadRequest,"File is too big")
		return
	}
	file,_,err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v",err)
		response.SendErrorResponse(w,http.StatusInternalServerError,"Internal error")
		return
	}
	data,err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v",err)
		response.SendErrorResponse(w,http.StatusInternalServerError,"Internal error")
		return

	}
	fname := p.ByName("vid-id")
	err = ioutil.WriteFile(defs.VIDEO_DIR+fname,data,0666)
	if err != nil {
		log.Printf("Write file error: %v",err)
		response.SendErrorResponse(w,http.StatusInternalServerError,"Internal error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w,"Upload successfully!")

}

func HqTestUploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

	filePath := defs.VIDEO_DIR + "upload.html"
	tpl,err := template.ParseFiles(filePath)
	if err != nil {
		response.SendErrorResponse(w,http.StatusNotFound,"Not Found")
		return
	}
	tpl.Execute(w,nil)

}
func HqHomeHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {


}