//easycode.go serves a webpage useful for showing a demo of Cloud Foundry
package main

import (
	"net/http"
	"html/template"
	"log"
)

type Location struct {
    Name string
}

func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}


func responseHandler(w http.ResponseWriter, r *http.Request) {
	location:=Location{Name: getName() }
	t,err:=template.ParseFiles("index.html")
	check("Parse template",err)
	t.Execute(w,location)
}

func getName() string{
	return "Philadelphia, PA"
}

func main() {
	http.HandleFunc("/",responseHandler)
	log.Fatal(http.ListenAndServe(":80",nil))
}