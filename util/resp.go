/*
@desc : Created by San on 2019/11/1 17:21
*/
package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data, omitempty"`
	Rows  interface{} `json:"rows, omitempty"`
	Total interface{} `json:"total, omitempty"`
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	w.Write([]byte(ret))
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

func RespOK(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}

	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}

	w.Write(ret)
}

func RespOKList(w http.ResponseWriter, lists interface{}, total interface{}) {
	RespList(w, 0, lists, total)
}
