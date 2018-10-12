package proxy

import (
	"net/http"
	"net/url"
	"net/http/httputil"
)
//proxy转发

func ProxyUploadHandler(w http.ResponseWriter,r *http.Request)  {
	//streamserver的地址
	u,_ := url.Parse("http://127.0.0.1:9000")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w,r)
}