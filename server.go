package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"
	err := os.WriteFile(fileName, p.Body, 0600)
	return err
}

func loadPage(fileName string) (*Page, error) {
	file := fileName + ".txt"
	body, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return &Page{Title: fileName, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
func main() {
	//	Execrsice 1
	//	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	//	p1.save()
	//	p2, _ := loadPage("TestPage")
	//	fmt.Println(string(p2.Body))

	//Exe 2:
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
