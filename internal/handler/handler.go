package handler

import (
	"fmt"
	"net/http"
	"strings"
)

type ApiHandler struct {
}

func (h ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Handler1 was used")
	if err := r.ParseForm(); err != nil {
		// если не заполнена, возвращаем код ошибки
		http.Error(w, "Bad auth", 401)
		return
	} else {
		path := strings.TrimPrefix(r.URL.Path, "/")
		fmt.Println(strings.Split(path, "/"))
	}
}

//type ApiHandler2 struct {
//}
//
//func (h ApiHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	//_, err := w.Write(h.Templ)
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	fmt.Println("Handler2 was used")
//}
