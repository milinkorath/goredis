package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"
)

var redisMap map[string]string

func main() {
	redisMap = make(map[string]string)
	http.HandleFunc("/", handleFunc)
	http.HandleFunc("/count/", countFunc)
	http.ListenAndServe(":8080", nil)
}
func handleFunc(w http.ResponseWriter, req *http.Request) {
	keyValue := req.URL.Path[1:]
	values := strings.Split(keyValue, "/")
	method := req.Method
	if method == "PUT" {
		if len(values) == 2 {
			redisMap[values[0]] = values[1]
			io.WriteString(w, "Key->"+values[0]+" Value->"+values[1]+" Saved\n")
		} else {
			io.WriteString(w, "Invalid Format-Format should be <key>/<value>")
		}
	} else if method == "GET" {
		if _, ok := redisMap[keyValue]; ok {
			io.WriteString(w, redisMap[keyValue]+"\n")
		} else {
			io.WriteString(w, "Key->"+keyValue+"Not Found\n")
		}

	} else if method == "DELETE" {
		if _, ok := redisMap[keyValue]; ok {
			delete(redisMap, keyValue)
			io.WriteString(w, "Key->"+keyValue+" deleted\n")
		} else {
			io.WriteString(w, "Key->"+keyValue+"Not Found\n")
		}
	}

}

func countFunc(w http.ResponseWriter, req *http.Request) {
	keyValue := req.URL.Path[1:]
	values := strings.Split(keyValue, "/")
	var count = 0
	for key := range redisMap {
		if len(values) == 2 {
			if strings.HasPrefix(key, values[1]) {
				count++
			}
		} else {
			count = len(redisMap)
			break
		}
	}
	io.WriteString(w, strconv.Itoa(count)+"\n")
}
