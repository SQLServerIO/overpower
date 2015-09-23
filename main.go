package main

import (
	"log"
	"mule/mylog"
	"net/http"
)

const (
	DATADIR  = "DATA/"
	SERVPORT = ":8080"
)

var Log = mylog.Err

func init() {
	mylog.SetErr(DATADIR + "errors.txt")
}

func main() {
	http.HandleFunc("/", gameMux)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("STATIC/"))))
	log.Println("STARTING SERVER AT", SERVPORT)
	http.ListenAndServe(SERVPORT, nil)
}