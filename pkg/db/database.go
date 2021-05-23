package db

import (
	"github.com/coreos/etcd/pkg/osutil"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"path"
)

func GetDB() *leveldb.DB {
	dbPath := path.Join(".", "data", "jabs.db")
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Printf("Could not init or find the database at %s", dbPath)
		osutil.Exit(-1)
	}
	return db
}
