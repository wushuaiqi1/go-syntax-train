package base

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Init 原生方式
func Init() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/hi", Hi)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

// Hello body是流式传递仅可读一次 第二次读不到不会报错
func Hello(w http.ResponseWriter, r *http.Request) {
	log.Println("hello execute...")
	data := r.Body
	bytes, err := io.ReadAll(data)
	if err != nil {
		return
	}
	//fmt.Println(string(json))
	w.WriteHeader(200)
	_, err = w.Write(bytes)
	if err != nil {
		return
	}
}

// Hi GetBody默认是nil
func Hi(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		// 向输出中
		fmt.Fprintf(w, "GetBody is nil")
	} else {
		fmt.Fprintf(w, "nis nil")
	}
}
