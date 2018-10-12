package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
	"fmt"
	"log"
	"html/template"
	"io/ioutil"
	"HqVideo/api/defs"
	"encoding/json"
	"HqVideo/api/response"
	"HqVideo/api/dbops"
	"HqVideo/api/session"
	"HqVideo/api/auth"
	"HqVideo/api/utils"
)

func HqGetUser()  {
	fmt.Println("dgdsgds")
}
func CreateUser(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {

	log.Printf("创建")
	res,_ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res,ubody); err != nil {
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	//插入用户
	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd); err != nil {
		response.SendErrorResponse(w,defs.ErrorDBError)
		return
	}
	//创建 session
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success:true,SessionId:id}
	if resp,err := json.Marshal(su); err != nil {
		response.SendErrorResponse(w,defs.ErrorInternalError)
		return
	}else {
		response.SendNormalResponse(w,string(resp),201)
	}
}
func HqGetUserInfo(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {

	if !auth.ValidateUser(w,r) {
		log.Printf("Unathor user\n")
		return
	}

	uname := p.ByName("username")
	u,err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("GetUserInfo Error :%s",err)
		response.SendErrorResponse(w,defs.ErrorDBError)
		return
	}
	ui := &defs.UserInfo{Id:u.Id}
	if resp,err := json.Marshal(ui);err != nil {
		response.SendErrorResponse(w,defs.ErrorInternalError)

	}else {
		response.SendNormalResponse(w,string(resp),200)
	}


	user := p.ByName("username")
	io.WriteString(w,user)
}
func Login(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {

	res,_:= ioutil.ReadAll(r.Body)
	log.Printf("Login %s",res)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res,ubody); err != nil {
		log.Printf("login json parse Error: %s",err)
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	//验证
	uname := p.ByName("username")
	log.Printf("login name =%s",uname)
	log.Printf("Login body name = %s",ubody.Username)
	if uname != ubody.Username {
		response.SendErrorResponse(w,defs.ErrorNotAuthUser)
		return
	}
	log.Printf("login ubody %s",ubody.Username)
	pwd,err := dbops.GetUserCredential(ubody.Username)
	log.Printf("login pwd = %s",pwd)
	log.Printf("login body pwd = %s",ubody.Pwd)
	if err != nil || len(pwd) ==0 || pwd != ubody.Pwd {
		response.SendErrorResponse(w,defs.ErrorNotAuthUser)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	si := &defs.SignedIn{Success:true,SessionId:id}
	if resp,err := json.Marshal(si); err != nil {
		response.SendErrorResponse(w,defs.ErrorInternalError)
	}else {
		response.SendNormalResponse(w,string(resp),200)
	}
	//user := p.ByName("username")
	//io.WriteString(w,user)
}
func AddNewVideo(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {

	if !auth.ValidateUser(w,r) {
		log.Printf("Unathor user\n")
		return
	}

	res,_ := ioutil.ReadAll(r.Body)
	nvbody :=defs.NewVideo{}
	if err := json.Unmarshal(res,nvbody); err != nil {
		log.Printf("Add video error:%s",err)
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	vi,err := dbops.AddNewVideo(nvbody.AuthorId,nvbody.Name)
	if err != nil {
		log.Printf("add new video db error：%s",err)
		response.SendErrorResponse(w,defs.ErrorDBError)
		return
	}

	if resp,err := json.Marshal(vi); err != nil {
		response.SendErrorResponse(w,defs.ErrorInternalError)
	}else {
		response.SendNormalResponse(w,string(resp),201)
	}
}
func ListAllVideo(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {
	if !auth.ValidateUser(w,r) {
		log.Printf("Unathor user\n")
		return
	}
	uname := p.ByName("username")
	vs,err := dbops.ListVideoInfo(uname,0,utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("error listAllVideo:%s",err)
		response.SendErrorResponse(w,defs.ErrorDBError)
		return
	}

	vsi := &defs.VideosInfo{Videos:vs}
	if resp,err := json.Marshal(vsi); err != nil {
		response.SendErrorResponse(w,defs.ErrorInternalError)
	}else {
		response.SendNormalResponse(w,string(resp),200)
	}
}
func DeleteVideo(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {
	if !auth.ValidateUser(w,r) {
		log.Printf("Unathor user\n")
		return
	}
	vid := p.ByName("vid-id")
	err := dbops.DeleteVideoInfo(vid)
	if err != nil {
		log.Printf("delete video info error:%s",err)
		response.SendErrorResponse(w,defs.ErrorDBError)
		return
	}

	response.SendNormalResponse(w,"success",204)

}

func PostComment(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {
	if !auth.ValidateUser(w,r) {
		log.Printf("Unathor user\n")
		return
	}
	reqBody,_ := ioutil.ReadAll(r.Body)
	cbody := &defs.NewComment{}
	if err := json.Unmarshal(reqBody,cbody); err != nil {
		log.Printf("post comment error:%s",err)
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	vid := p.ByName("vid-id")
	if err := dbops.AddNewComments(vid,cbody.AuthorId,cbody.Content); err != nil {
		log.Printf("add new comment db error:%s",err)
		response.SendErrorResponse(w,defs.ErrorDBError)

	}else {
		response.SendNormalResponse(w,"success",201)
	}
}
func ShowComments(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {
	if !auth.ValidateUser(w,r) {
		log.Printf("Unathor user\n")
		return
	}
	vid := p.ByName("vid-id")
	cm,err := dbops.ListComments(vid,0,utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("show comments error:%s",err)
		response.SendErrorResponse(w,defs.ErrorDBError)
		return
	}
	cms := &defs.Comments{cm}
	if resp,err := json.Marshal(cms); err != nil {
		response.SendErrorResponse(w,defs.ErrorInternalError)
	}else {
		response.SendNormalResponse(w,string(resp),200)
	}

}
func HomePage(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {

	// os.Getwd() 获取当前项目root目录
	/*
	tplPath,_ := os.Getwd()
	tplPath = tplPath + "/api/templates/index.html"
	fmt.Println(tplPath)
	*/

	//根据相对绝对路径获取文件绝对路径
	//tplPath,_ := filepath.Abs("./api/templates/index.html")

	tplPath := http.Dir("./api/templates/index.html")

	fmt.Println(tplPath)
	tpl,err := template.ParseFiles(string(tplPath))
	fmt.Println(tpl.Name())

	if err != nil {
		log.Println("err==",err)
		return
	}

	//tpl.Execute(w,"视频大家")
	home := Home{"首页","各种视频分类"}
	err = tpl.ExecuteTemplate(w,"index.html",home)
	if err != nil {
		fmt.Println(err)
	}

}

type Home struct {
	Title string
	Content string
}
