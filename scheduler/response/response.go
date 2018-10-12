package response

import (
	"net/http"
	"io"
)

func SendResponse(w http.ResponseWriter,sc int,msg string)  {
	w.WriteHeader(sc)
	io.WriteString(w,msg)

}