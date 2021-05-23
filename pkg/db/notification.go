package db

import (
	"github.com/katta/jabfinder/pkg/models"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"time"
)

func Register(sessions []models.FlatSession) []models.FlatSession {
	db := GetDB()
	defer db.Close()

	newSessions := []models.FlatSession{}
	for _, session := range sessions {
		_, err := db.Get([]byte(session.SessionId), nil)

		if err == leveldb.ErrNotFound {
			err = db.Put([]byte(session.SessionId), []byte(time.Now().String()), nil)
			if err != nil {
				log.Printf("Error while writing session %s to the db", session.SessionId)
			} else {
				newSessions = append(newSessions, session)
			}
		}
	}

	return newSessions
}
