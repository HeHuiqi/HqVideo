package response

import (
	"net/http"
	"io"
)

func SendErrorResponse(w http.ResponseWriter,sc int,errMsg string)  {
	w.WriteHeader(sc)
	io.WriteString(w,errMsg)

}