package main

import (
	"flag"
	"log"

	"github.com/matryer/filedb"
	"github.com/tanagaku/sample_go_learning/chapter8/backup"
)

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

	_ = &backup.Monitor{
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
	/**col*/ _, err = db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}

}
