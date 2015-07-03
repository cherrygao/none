package platform

import (
	"io"
	"net/http"
)

func NewDev() (http.ServeMux, error) {
	d := newDev()
	mux := http.NewServeMux()
	mux.HandleFunc("/push", d.push)
	mux.HandleFunc("/pushall", d.pushAll)
}

type dev struct {
}

func newDev() *dev {

}

func (self *dev) push(w http.ResponseWriter, req *http.Request) {

	m := req.URL.Query()
	mode := m.Get("mode")
	users := m.Get("users")
}

func (self *dev) pushAll(w http.ResponseWriter, req *http.Request) {

	m := req.URL.Query()
	mode := m.Get("mode")
}

func (self *dev) sendMail(mode string, users []string, message io.Reader) {

}
