package webserver

import (
	"fmt"
	"log"
	// "io/ioutil"
	"net/http"
	"go-practice/gowiki"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "Hello there, I love %s!", r.URL.Path[1:])
}

func ViewHandler(w http.ResponseWriter, r *http.Request)  {
	title := r.URL.Path[len("/view/"):]
	p, _ := gowiki.LoadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func Execute(port string)  {
	address := "http://localhost" + port
	fmt.Printf("\n\n Starting webserver at: '%v' \n\n", address)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}

func ExecuteViewHandler()  {
	http.HandleFunc("/view/", ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}