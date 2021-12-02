package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/bitly/go-nsq"
	"gopkg.in/mgo.v2"
)

var fatalErr error
var countsLock sync.Mutex
var counts map[string]int

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalErr = e
}

func main() {
	defer func() {
		if fatalErr != nil {
			os.Exit(1)
		}
	}()

	log.Println("データベースに接続します")
	db, err := mgo.Dial("localhost")
	if err != nil {
		fatal(err)
		db.Close()
	}
	defer func() {
		log.Println("データベース接続を閉じます")
		db.Close()
	}()
	db.DB("balottos").C("polls")

	log.Println("NSQに接続します...")
	q, err := nsq.NewConsumer("votes", "counter", nsq.NewConfig())
	if err != nil {
		fatal(err)
		return
	}

	q.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		countsLock.Lock()
			defer countsLock.Unlock()
			if counts == nil{
				counts=make(map[string]int])
			}
			vote := string(m.Body)
			counts[vote]++
			return nil
	}))
}
