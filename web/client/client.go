package client
//api透传
import (
	"net/http"
	"HqVideo/web/defs"
	"log"
	"io/ioutil"
	"io"
	"bytes"
)

var httpClient *http.Client

func init()  {
	httpClient = &http.Client{}
}
func Request( b *defs.ApiBody,w http.ResponseWriter,r *http.Request)  {

	var resp *http.Response
	var err error
	switch b.Method {
	case http.MethodGet:
		req,_ := http.NewRequest("GET",b.Url,nil)
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			normalResponse(w,resp)
			return
		}
		normalResponse(w,resp)
	case http.MethodPost:
		req,_ := http.NewRequest("POST",b.Url,bytes.NewBuffer([]byte(b.ReqBody)))
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			normalResponse(w,resp)
			return
		}
		normalResponse(w,resp)
	case http.MethodDelete:
		req,_ := http.NewRequest("DELETE",b.Url,nil)
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			normalResponse(w,resp)
			return
		}
		normalResponse(w,resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w,"Bat Request")
		return
	}


}

func normalResponse(w http.ResponseWriter,rs *http.Response)  {

	res,err := ioutil.ReadAll(rs.Body)
	if err != nil {

		w.WriteHeader(500)
		io.WriteString(w,string(res))
		return
	}

	w.WriteHeader(rs.StatusCode)
	io.WriteString(w,string(res))

}