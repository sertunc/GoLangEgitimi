/*
Kullanılabilecek kütüphaneler

https://golang.org/pkg/net/http/
https://github.com/gorilla/mux

*/

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("Hello World"))
	})

	log.Println("Serving...")
	http.ListenAndServe("localhost:3000", nil)
}
