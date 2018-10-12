package auth

import (
	"net/http"
	"HqVideo/api/session"
	"HqVideo/api/response"
	"HqVideo/api/defs"
)

var HEADER_FIELD_SESSION  = "X-Session-Id"
var HEADER_FIELD_UNAME  = "X-User-Name"

func ValidateUserSession(r *http.Request) bool {

	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	uname,ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME,uname)
	return true
}
// IAM 用户权限管理
// SSO
// Rbac
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {

	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		response.SendErrorResponse(w,defs.ErrorNotAuthUser)
		return  false
	}
	return true
}
