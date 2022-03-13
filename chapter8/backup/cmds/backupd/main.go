package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"

	"github.com/matryer/filedb"
	"github.com/tanagaku/sample_go_learning/chapter8/backup"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()
	var (
		//	interval = flag.Int("interval", 10, "チェックの感覚(秒単位")
		archive = flag.String("archive", "archive", "アーカイブの保存先")
		dbpath  = flag.String("db", "./db", "filedbデータベースへのパス")
	)

	m := &backup.Monitor{
		Destination: *archive,
		Archiver:    backup.ZIP,
		//Path:        make(map[string]string{}),
	}
	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()
	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}

	var path path
	col.ForEach(func(_ int, data []byte) bool {
		if err := json.Unmarshal(data, &path); err != nil {
			fatalErr = err
			return true
		}
		m.Paths[path.Path] = path.Hash
		return false //処理を中断します
	})
	if fatalErr != nil {
		return
	}
	if len(m.Paths) < 1 {
		fatalErr = errors.New("パスがりません。backupツールを使って追加してください")
		return
	}
}
