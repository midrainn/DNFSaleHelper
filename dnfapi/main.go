package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./httpsession"
	"./origin"
	"./sqlservice"
)

//检测是否登录，这个是服务器内部函数，不是API
func isLogin(w http.ResponseWriter, r *http.Request) bool {
	session := httpsession.SessionManager.GetSession(w, r)
	if session.GetAttribute("userid") == nil {
		return false
	}

	return true
}

//检测是否登录，这是暴露的API
func isLoginAPI(w http.ResponseWriter, r *http.Request) {
	if origin.IsAllowOrigin(r.Header.Get("Origin")) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if !isLogin(w, r) {
			fmt.Fprint(w, "CODE2001") //尚未登录
		} else {
			fmt.Fprint(w, "CODE1001") //已经登录

		}
	} else {
		fmt.Println(r.Header.Get("Origin") + "跨域被阻止")
	}
}
func saveSave(w http.ResponseWriter, r *http.Request) {
	if origin.IsAllowOrigin(r.Header.Get("Origin")) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if !isLogin(w, r) {
			fmt.Fprint(w, "CODE2001") //尚未登录
		} else {
			bodystr := make([]byte, r.ContentLength)
			r.Body.Read(bodystr)
			savemap := make(map[string]string)
			json.Unmarshal(bodystr, &savemap)
			session := httpsession.SessionManager.GetSession(w, r)

			if sqlservice.SaveSave(session.GetAttribute("userid").(string), savemap["saves"]) {
				fmt.Fprint(w, "CODE1002") //提交成功
			} else {
				fmt.Fprint(w, "CODE2003") //提交失败
			}

		}
	} else {
		fmt.Println(r.Header.Get("Origin") + "跨域被阻止")
	}

}

func loadSave(w http.ResponseWriter, r *http.Request) {
	if origin.IsAllowOrigin(r.Header.Get("Origin")) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if !isLogin(w, r) {
			fmt.Fprint(w, "CODE2001") //尚未登录
		} else {

			session := httpsession.SessionManager.GetSession(w, r)
			savs := sqlservice.GetSave(session.GetAttribute("userid").(string))
			fmt.Fprint(w, savs)
		}
	} else {
		fmt.Println(r.Header.Get("Origin") + "跨域被阻止")
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	if origin.IsAllowOrigin(r.Header.Get("Origin")) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		bodystr := make([]byte, r.ContentLength)
		r.Body.Read(bodystr)
		usermap := make(map[string]string)
		json.Unmarshal(bodystr, &usermap)
		if uid := sqlservice.IsUser(usermap["username"], usermap["password"]); uid != "" {
			session := httpsession.SessionManager.GetSession(w, r)
			session.SetAttribute("userid", uid)
			fmt.Fprint(w, "CODE1000") //登录成功
		} else {
			fmt.Fprint(w, "CODE2002") //账号或密码错误
		}

	} else {
		fmt.Println(r.Header.Get("Origin") + "跨域被阻止")
	}

}

func main() {
	http.HandleFunc("/api/saveislogin", isLoginAPI)
	http.HandleFunc("/api/dnflogin", login)
	http.HandleFunc("/api/saveSale", saveSave)
	http.HandleFunc("/api/loadSalesaves", loadSave)
	fmt.Println("服务器启动完毕")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
