package response

import (
	"net/http"
	"HqVideo/api/defs"
	"io"
	"encoding/json"
)

func SendErrorResponse(w http.ResponseWriter,errResp defs.ErrorResponse)  {

	w.WriteHeader(errResp.HttpSC)
	respStr,_ := json.Marshal(&errResp.Error)
	io.WriteString(w,string(respStr))

}

func SendNormalResponse(w http.ResponseWriter,resp string,sc int)  {

	w.WriteHeader(sc)
	io.WriteString(w,resp)
}