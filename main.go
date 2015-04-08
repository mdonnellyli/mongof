package main

import (
	"encoding/json"

	"github.com/pjvds/tidy"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	log := tidy.GetLogger()
	session, err := mgo.Dial("localhost")

	if err != nil {
		log.With("error", err).With("url", "localhost").Fatal("failed to connect")
	}

	localdb := session.DB("local")

	verifyOplog := func() (bool, error) {
		collections, err := localdb.CollectionNames()

		if err != nil {
			return false, err
		}

		for _, name := range collections {
			log.With("name", name).Debug("collection found")

			if name == "oplog.$main" {
				return true, nil
			}
		}
		return false, nil
	}

	if exists, err := verifyOplog(); !exists {
		if err != nil {
			log.With("error", err).Error("failed to verify oplog")
		}
		log.Fatalf("oplog is missing, are you running a replicate set or standalone master?")
	}

	oplog := localdb.C("oplog.$main")
	cursor := oplog.Find(nil).Tail(-1)

	var document bson.M
	for cursor.Next(&document) {
		j, _ := json.MarshalIndent(document, "", "\t")
		println(string(j))
	}

	if err := cursor.Err(); err != nil {
		log.With("error", err).Fatalf("tail query failed")
	}
}