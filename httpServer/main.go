package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting my http server!")
	muxServer := http.NewServeMux()
	muxServer.HandleFunc("/", requestResponseHandler)
	muxServer.HandleFunc("/version", versionHandler)
	muxServer.HandleFunc("/httpInfo", httpInfoHandler)
	muxServer.HandleFunc("/healthz", healthzHandler)
	fmt.Println("Starting my http server")
	err := http.ListenAndServe(":80", muxServer)
	if err != nil {
		log.Fatal(err)
	}

}

func requestResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering requestResponseHandler!")
	RequestHeaders := r.Header
	for header := range RequestHeaders {
		headerValues := RequestHeaders[header]
		for i, _ := range headerValues {
			headerValues[i] = strings.TrimSpace(headerValues[i])
		}
		w.Header().Set(header, strings.Join(headerValues, ","))
	}

	io.WriteString(w, "header数据：\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("header-key:%s=header-value:%s\n", k, v))
	}
}

func httpInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering httpInfoHandler!")
	remoteAddr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}
	println("Client-ip=" + remoteAddr)
	println("Client response code:" + strconv.Itoa(http.StatusOK))
	io.WriteString(w, "Visit the httpInfo ok")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering versionHandler!")
	os.Setenv("VERSION", "GOVERSION = go1.19.3")
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	io.WriteString(w, "Visit the version ok!")

}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering healthzHandler!")
	w.WriteHeader(200)
	io.WriteString(w, "Visit the healthz ok!")
}
