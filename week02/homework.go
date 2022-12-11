package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "runtime"
)

func main() {
    startServer()
}

func startServer()  {
    initRouter()

    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func initRouter()  {
    http.HandleFunc("/", handleRequest)
    http.HandleFunc("/healthz", handleHealthz)
}

func handleRequest(w http.ResponseWriter, r * http.Request){

    rHeader := r.Header

    //1.接收客户端 request，并将 request 中带的 header 写入 response header
    for k, v := range rHeader {
        w.Header().Add(k, sliceToString(v))
    }

    //2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
    w.Header().Add("VERSION",runtime.Version())

    //3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
    log.Printf("IP:%s, Http Code:%s",r.Host,w.Header().Get("statusCode"))
}

//4.当访问 localhost/healthz 时，应返回200
func handleHealthz(w http.ResponseWriter, r * http.Request){

    handleRequest(w,r)
    //4.当访问 localhost/healthz 时，应返回200
    w.WriteHeader(200)
    io.WriteString(w, "200")
}


func sliceToString(s []string) string {
    if len(s) == 0 {
        return ""
    }
    sStr := fmt.Sprintf("%v",s)
    return sStr[2:len(sStr)-1]
}