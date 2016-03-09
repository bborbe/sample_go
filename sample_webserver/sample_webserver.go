package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s handle request", time.Now().Format("2006-01-02T15:04:05"))
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, `<doctype html>`)
	fmt.Fprintf(res, `<html>`)
	fmt.Fprintf(res, `<head>`)
	fmt.Fprintf(res, `<title>Hello World</title>`)
	fmt.Fprintf(res, `</head>`)
	fmt.Fprintf(res, `<body>`)
	fmt.Fprintf(res, `<h1>Hello World</h1>`)
	fmt.Fprintf(res, time.Now().Format("2006-01-02T15:04:05"))
	fmt.Fprintf(res, `</body>`)
	fmt.Fprintf(res, `</html>`)
}

func main() {
	log.Printf("%s webserver started", time.Now().Format("2006-01-02T15:04:05"))
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
