package main

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"html/template"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var globalSessions *session.Manager

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello 123456!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		// t, _ := template.ParseFiles("login.gtpl")
		// log.Println(t.Execute(w, nil))
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
		sess.Set("token", token)
	} else {
		r.ParseForm() //解析参数，默认是不会解析的
		//请求的是登录数据，那么执行登录的逻辑判断
		if len(r.Form["username"][0]) == 0 {
			//为空的处理
			fmt.Fprintf(w, "username is null")
			return
		}

		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Fprintf(w, "email format error")
			return
		}

		token := r.Form.Get("token")
		if token != "" && sess.Get("token") != token {
			fmt.Fprintln(w, token)
			fmt.Fprintf(w, "token error")
			return
		} else {
			sess.Delete("token")
		}

		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		// fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		// fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
		// fmt.Fprintf(w, strings.Join(r.Form["username"], ""))
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
