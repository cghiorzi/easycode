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
	t,err:=template.ParseFiles("templates/index.tmpl")
	check("Parse template",err)
	t.Execute(w,location)
}

func getName() string{
	return "Philadelphia, PA"
}

func main() {
	http.HandleFunc("/",responseHandler)
	http.Handle("/images/",http.StripPrefix("/images/",http.FileServer(http.Dir("/home/vcap/app/images"))))
	http.Handle("/css/",http.StripPrefix("/css/",http.FileServer(http.Dir("/home/vcap/app/css"))))
	http.Handle("/fonts/",http.StripPrefix("/fonts/",http.FileServer(http.Dir("/home/vcap/app/fonts"))))
	http.Handle("/js/",http.StripPrefix("/js/",http.FileServer(http.Dir("/home/vcap/app/js"))))
	log.Fatal(http.ListenAndServe(":8080",nil))
}