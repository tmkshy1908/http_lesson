package main

import (
	"log"
	"net/http"
)

type templateHandler struct {
	once sync.Onse
	filename string
	templ *template.Template
}

func (t *templateHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func(){
		t.templ = template.Must(template.ParseFiles(filepath.Join("temolates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		
		`))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"golaing.org/x/net/http2"
// )

// type MyHandler struct{}

// func (h *MyHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }

// func main() {
// 	handler := MyHandler{}
// 	server := http.Server{
// 		Addr:    "127.0.0.1:8080",
// 		Handler: &handler,
// 	}
// 	http2.ConfigusreSeveTLS("cert.pem", "key.pem")
// 	server.ListenAndServeTLS("cert.pem", "key.pem")
// }
