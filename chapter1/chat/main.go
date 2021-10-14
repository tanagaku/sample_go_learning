package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

//templ1は１つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServerHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	if err := t.templ.Execute(w, nil); err != nil {
		log.Fatal("templateHandleError", err)
	}
}

func main() {
	//ルート
	http.Handle("/", &templateHandler{filename: "chat.html"})

	//webサーバーを開始します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
