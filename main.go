/*
@desc : Created by San on 2019/10/29 01:44
*/
package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello World")
	})

	http.ListenAndServe(":8080", nil)
}
