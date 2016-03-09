package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(res http.ResponseWriter, req *http.Request) {
	date := time.Now().Format("2006-01-02T15:04:05")
	fmt.Printf("handle request %s\n", date)
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, `<doctype html>`)
	fmt.Fprintf(res, `<html>`)
	fmt.Fprintf(res, `<head>`)
	fmt.Fprintf(res, `<title>Hello World</title>`)
	fmt.Fprintf(res, `</head>`)
	fmt.Fprintf(res, `<body>`)
	fmt.Fprintf(res, `<h1>Hello World</h1>`)
	fmt.Fprintf(res, date)
	fmt.Fprintf(res, `</body>`)
	fmt.Fprintf(res, `</html>`)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
