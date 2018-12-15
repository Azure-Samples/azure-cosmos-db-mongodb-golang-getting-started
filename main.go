package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/Azure/go-autorest/autorest/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	database string
	password string
)

func init() {
	database = utils.GetEnvVarOrExit("AZURE_DATABASE")
	password = utils.GetEnvVarOrExit("AZURE_DATABASE_PASSWORD")
}

// Package represents a document in the collection
type Package struct {
	Id            bson.ObjectId `bson:"_id,omitempty"`
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func main() {
	// DialInfo holds options for establishing a session with Azure Cosmos DB for MongoDB API account.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s.documents.azure.com:10255", database)}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: database, // It can be anything
		Username: database, // Username
		Password: password, // PASSWORD
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	// Create a session which maintains a pool of socket connections
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("Can't connect, go error %v\n", err)
		os.Exit(1)
	}

	defer session.Close()

	// SetSafe changes the session safety mode.
	// If the safe parameter is nil, the session is put in unsafe mode, and writes become fire-and-forget,
	// without error checking. The unsafe mode is faster since operations won't hold on waiting for a confirmation.
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode.
	session.SetSafe(&mgo.Safe{})

	// get collection
	collection := session.DB(database).C("package")

	// insert Document in collection
	err = collection.Insert(&Package{
		FullName:      "react",
		Description:   "A framework for building native apps with React.",
		ForksCount:    11392,
		StarsCount:    48794,
		LastUpdatedBy: "shergin",
	})

	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}

	// Get Document from collection
	result := Package{}
	err = collection.Find(bson.M{"fullname": "react"}).One(&result)
	if err != nil {
		log.Fatal("Error finding record: ", err)
		return
	}

	fmt.Println("Description:", result.Description)

	// update document
	updateQuery := bson.M{"_id": result.Id}
	change := bson.M{"$set": bson.M{"fullname": "react-native"}}
	err = collection.Update(updateQuery, change)
	if err != nil {
		log.Fatal("Error updating record: ", err)
		return
	}

	// delete document
	err = collection.Remove(updateQuery)
	if err != nil {
		log.Fatal("Error deleting record: ", err)
		return
	}
}
