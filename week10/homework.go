package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

// 创建一个简单呃 gauge 指标。
var requestDurations = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "http_request_duration_seconds",
	Help:    "A histogram of the HTTP request durations in seconds.",
	Buckets: prometheus.ExponentialBuckets(0.0001, 2, 15),
})

func main() {
	fmt.Println("server is starting......")
	startServer()
}

func startServer() {
	initRouter()

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initRouter() {
	//添加监控指标

	//注册metrics
	http.Handle("/metrics", promhttp.Handler())
	registryMetrics()

	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/healthz", handleHealthz)
}

func registryMetrics() {
	// 创建一个自定义的注册表
	registry := prometheus.NewRegistry()

	//注册直方图
	registry.MustRegister(requestDurations)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 启动一个计时器
	timer := prometheus.NewTimer(requestDurations)

	//为 HTTPServer 添加 0-2 秒的随机延时；
	time.Sleep(time.Duration(rand.Intn(2*1000)) * time.Millisecond)

	rHeader := r.Header
	//1.接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range rHeader {
		w.Header().Add(k, sliceToString(v))
	}

	//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Add("VERSION", runtime.Version())

	//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	log.Printf("client IP:%s, reponse Code:%d", r.RemoteAddr, 200)

	// 停止计时器并观察其持续时间，将其放进 requestDurations 的直方图指标中去
	timer.ObserveDuration()
}

// 4.当访问 localhost/healthz 时，应返回200
func handleHealthz(w http.ResponseWriter, r *http.Request) {

	handleRequest(w, r)
	//4.当访问 localhost/healthz 时，应返回200
	w.WriteHeader(200)
	_, err := io.WriteString(w, "200")
	if err != nil {
		log.Println(err)
	}
}

func sliceToString(s []string) string {
	if len(s) == 0 {
		return ""
	}
	sStr := fmt.Sprintf("%v", s)
	return sStr[2 : len(sStr)-1]
}
