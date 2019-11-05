/*
@desc : Created by San on 2019/10/29 01:44
*/
package main

import (
	"chatting/ctrl"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, v := range tpl.Templates() {
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			err = tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}

func main() {
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	RegisterView()
	port := "8080"
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
