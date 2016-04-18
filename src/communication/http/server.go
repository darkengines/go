package http

import "net/http"
import "service"
import _ "fmt"
import _ "reflect"

type Server struct {
}

func NewServer() *Server {
	server := new(Server)
	return server
}

func handler(w http.ResponseWriter, r *http.Request) {
	context := HttpContext{Request: r, Response: w}
	err := service.Process(context)
	if (err != nil) {
		panic(err)
	}
}

func (this *Server) Start() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func init() {

}
