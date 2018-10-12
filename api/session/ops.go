package session

import (
	"sync"
	"time"
	"HqVideo/api/dbops"
	"HqVideo/api/defs"
	"HqVideo/api/utils"
)
//同步map保证线程安全
var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

func nowInmilli() int64  {

	return time.Now().UnixNano()/1000000
}

func DeleteExpiredSession(sid string)  {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionFromDB()  {

	r,err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(key, value interface{}) bool {

		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key,ss)
		return true
	})

}

func GenerateNewSessionId(uname string) string  {

	id,_ := utils.NewUUID()
	ctime := nowInmilli()
	ttl := ctime + 30 * 60 * 100//30分钟后过期
	ss := &defs.SimpleSession{Username:uname,TTL:ttl}
	sessionMap.Store(id,ss)
	dbops.InsertSession(id,ttl,uname)
	return id
}

func IsSessionExpired(sid string) (string,bool)  {

	ss , ok := sessionMap.Load(sid)
	if ok {
		ct := nowInmilli()
		if ss.(*defs.SimpleSession).TTL <ct  {
			DeleteExpiredSession(sid)
			return "",true
		}
		return ss.(*defs.SimpleSession).Username,false
	}
	return "",true
}