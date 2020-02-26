package httpsession

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

//SessionManager SESSION管理器,用来获取SESSION
var SessionManager sessionManager

//SESSION超时时间,单位秒
const sessionOverTimeTrick = 3600

//SESSION的回收间隔，单位秒
const sessionGCTimeTrick = 60

//Session Session的实体
type Session struct {
	id        string
	attribute map[string]interface{}
	overtime  time.Time
}

//SESSION管理器，用于保存SEESION,获取SEESION
type sessionManager struct {
	sessions map[string]Session
}

//初始化模块,初始化管理器，并启动SESSIONGC
func init() {
	defer fmt.Println("SESSION模块初始化完毕")
	SessionManager = sessionManager{
		sessions: make(map[string]Session),
	}
	go sessionGC()
}

//回收过期SESSION
func sessionGC() {
	tircker := time.NewTicker(time.Second * sessionGCTimeTrick)
	for {
		_ = <-tircker.C
		for k, v := range SessionManager.sessions {
			if v.overtime.Before(time.Now()) {
				delete(SessionManager.sessions, k)
			}
		}
	}
}

//给SESSION续期
func (session Session) resetTime() Session {
	overtime, _ := time.ParseDuration(strconv.Itoa(sessionOverTimeTrick) + "s")
	session.overtime = time.Now().Add(overtime)
	return session
}

//SetAttribute 设置一个值
func (session Session) SetAttribute(key string, value interface{}) {
	if session.attribute == nil {
		session.attribute = make(map[string]interface{})
	}
	session.attribute[key] = value
}

//GetAttribute 从seesion中获取值
func (session Session) GetAttribute(key string) interface{} {
	if session.attribute == nil {
		return nil
	}
	v, flag := session.attribute[key]
	if flag {
		return v
	}
	return nil
}

//获取SESSION，如果没有SESSION会创建一个SESSION
func (manager sessionManager) GetSession(response http.ResponseWriter, request *http.Request) *Session {

	if manager.sessions == nil {
		manager.sessions = make(map[string]Session)
	}
	key, err := request.Cookie("SESSIONID")
	if err != nil || len(key.Value) == 0 {
		newid := getNewUUID()
		overtime, _ := time.ParseDuration(strconv.Itoa(sessionOverTimeTrick) + "s")
		newsession := Session{
			id:        newid,
			attribute: make(map[string]interface{}),
			overtime:  time.Now().Add(overtime),
		}
		manager.sessions[newid] = newsession
		http.SetCookie(response, &http.Cookie{Name: "SESSIONID", Value: newid})
		return &newsession
	}
	v, flag := manager.sessions[key.Value]
	if !flag {

		newid := getNewUUID()
		overtime, _ := time.ParseDuration(strconv.Itoa(sessionOverTimeTrick) + "s")
		newsession := Session{
			id:        newid,
			attribute: make(map[string]interface{}),
			overtime:  time.Now().Add(overtime),
		}
		v = newsession
		manager.sessions[newid] = newsession
		http.SetCookie(response, &http.Cookie{Name: "SESSIONID", Value: newid})
	} else {
		manager.sessions[key.Value] = manager.sessions[key.Value].resetTime()
	}
	return &v

}

//获取一个新的UUID，用于SESSION_ID
func getNewUUID() string {
	newid := ""

	for flag := true; flag == true; {
		newid = uuid.Must(uuid.NewV4()).String()
		flag = false
		for _, v := range SessionManager.sessions {
			if v.id == newid {
				flag = true
				break
			}
		}
	}

	return newid
}
