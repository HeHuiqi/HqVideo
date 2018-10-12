package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"fmt"
	"encoding/json"
	"HqVideo/web/defs"
	"io"
	"io/ioutil"
	"HqVideo/web/client"
	"HqVideo/web/proxy"
)

type HomePage struct {
	Title string
}
type UserPage struct {
	Name string
} 

func HomeHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {


	cname , err1 := r.Cookie("username")
	sid,err2 := r.Cookie("session")
	fmt.Println("canme--",cname,"sid",sid)

	if err1 != nil || err2 != nil {
		home := &HomePage{Title:"视频网站首页"}
		tpl,err := template.ParseFiles("./web/templates/home.html")
		if err != nil {

			log.Printf("Parseing template home.html error: %v",err)
			return
		}
		//tpl.ExecuteTemplate(w,"home.html",home)
		tpl.Execute(w,home)
		return
	}
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w,r,"/userhome",http.StatusFound)
		return
	}

}

func UserHomeHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

	name , err1 := r.Cookie("username")
	_,err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		http.Redirect(w,r,"/",http.StatusFound)
		return

	}
	fname := r.FormValue("username")

	var user *UserPage

	if len(name.Value) != 0 {
		user = &UserPage{Name:name.Value}
	}else if len(fname) != 0 {
		user = &UserPage{Name:fname}
	}
	tpl,err := template.ParseFiles("./web/templates/username.html")
	if err != nil {

		log.Printf("Parseing template username.html error: %v",err)
		return
	}
	tpl.Execute(w,user)
}

func ApiHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

	if r.Method != http.MethodPost {
		
		re, _ := json.Marshal(defs.ErrorRequestNotRecognized)
		
		io.WriteString(w,string(re))
		return 
	}
	res,_ := ioutil.ReadAll(r.Body)
	apibody := &defs.ApiBody{}
	if err := json.Unmarshal(res,apibody);err != nil {
		res,_ := json.Marshal(defs.ErrorRequestBodyParseFailed)
		io.WriteString(w,string(res))
		return
	}
	client.Request(apibody,w,r)
	defer r.Body.Close()
	
}
//proxy转发
func ProxyUploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

	proxy.ProxyUploadHandler(w,r)

}