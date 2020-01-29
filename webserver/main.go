package webserver

import (
	"fmt"
	"log"
	// "io/ioutil"
	"net/http"
	"html/template"
	"go-practice/gowiki"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "Hello there, I love %s!", r.URL.Path[1:])
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *gowiki.Page)  {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)	
}

func ViewHandler(w http.ResponseWriter, r *http.Request)  {
	title := r.URL.Path[len("/view/"):]
	p, _ := gowiki.LoadPage(title)
	renderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, err := gowiki.LoadPage(title)
	if err != nil {
		p = &gowiki.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// func SaveHandler()  {
	
// }

// func Execute(port string)  {
// 	address := "http://localhost" + port
// 	fmt.Printf("\n\n Starting webserver at: '%v' \n\n", address)
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(port, nil))
// }

// func ExecuteViewHandler()  {
// 	http.HandleFunc("/view/", ViewHandler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func ExecuteWiki(port string)  {
	address := "http://localhost" + port
	fmt.Printf("\n\n Starting webserver at: '%v' \n\n", address)
	http.HandleFunc("/view/", ViewHandler)
	http.HandleFunc("/edit/", EditHandler)
	// http.HandleFunc("/save/", SaveHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}

func RunTest()  {
	Run()
}