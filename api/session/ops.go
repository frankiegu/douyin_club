package session

import (
	"time"
	"sync"
	"github.com/Yq2/douyin_club/api/defs"
	"github.com/Yq2/douyin_club/api/dbops"
	"github.com/Yq2/douyin_club/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64{
	return time.Now().UnixNano()/1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool{
		//对v进行类型断言，v是interface接口类型，所以可以直接断言
		ss := v.(*defs.SimpleSession)
		//将符合类型的数据存储到全局session map中
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	//session过期时间为30分钟
	ttl := ct + 60 *60 * 60 * 1000// Severside session valid time: 2h

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	ct := nowInMilli()
	if ok {
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	} else {
		ss, err := dbops.RetrieveSession(sid)
		if err != nil || ss == nil {
			return "", true
		}

		if ss.TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		sessionMap.Store(sid, ss)
		return ss.Username, false
	}

	return "", true
}
