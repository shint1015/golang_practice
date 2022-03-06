package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

//ファイルを作成、書き込み
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//ページの読み込み
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

var templates = template.Must(template.ParseFiles("edit.html, view.html"))


func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//t, _ := template.ParseFiles(tmpl + ".html")
	//t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, _ := loadPage(title)

	//少量のhtmlを出力する時
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	//templateHTMLへ書き加える時
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//formから値を取得
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		//errorcode
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//リダイレクト処理
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")


func  makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil{
			http.NotFound(w,r)
			return
		}
		fn(w,r,m[2])
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	title := "miyata"
	p, _ := loadPage(title)

	//少量のhtmlを出力する時
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	//templateHTMLへ書き加える時
	//renderTemplate(w, "view", p)
}

func main() {
	http.HandleFunc("/miyata/", testHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatalln(http.ListenAndServe(":6060", nil))
}
