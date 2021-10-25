package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

//templ1は１つのテンプレートを表します
type templateHandler struct {
	once     sync.Once //複数から呼び出されても関数が1度しか実行されないことを保証する
	filename string
	templ    *template.Template
}

// ServerHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	if err := t.templ.Execute(w, r); err != nil {
		log.Fatal("templateHandleError", err)
	}
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() //フラグを解釈
	r := newRoom()
	//ルート
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	//チャットルームを開始
	go r.run()
	log.Println("webサーバーを開始します。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
