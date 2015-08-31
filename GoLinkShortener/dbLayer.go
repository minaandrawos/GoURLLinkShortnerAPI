/*
	This is the database layer used in the program. The GoLinkShortner API uses a mongo database as a backend.
	The database layer uses the mgo library.
	It follows a common connection pattern where a main session is created then other sessions are created by copying the information of the main session while utilizing a different socket from a socket pool
 */
package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"fmt"
)
//this is subject to change based on the connection parameters, could also be configurable
const CONNECTIONSTRING = "mongodb://127.0.0.1"

type mongoDocument struct{
	Id bson.ObjectId    `bson:"_id"`
	ShortUrl string     `bson:"shorturl"`
	LongUrl string      `bson:"longurl"`
}

type MongoConnection struct{
	originalSession *mgo.Session
}

func NewDBConnection() (conn *MongoConnection) {
	conn = new(MongoConnection)
	conn.createLocalConnection()
	return
}

func (c *MongoConnection) createLocalConnection() (err error){
    fmt.Println("Connecting to local mongo server....")
	c.originalSession, err = mgo.Dial(CONNECTIONSTRING)
	if err == nil {
		fmt.Println("Connection established to mongo server")
		urlcollection := c.originalSession.DB("LinkShortnerDB").C("UrlCollection")
		if urlcollection == nil {
			err = errors.New("Collection could not be created, maybe need to create it manually")
		}
		//This will create a unique index to ensure that there won't be duplicate shorturls in the database.
		index := mgo.Index{
			Key: []string{"$text:shorturl"},
			Unique: true,
			DropDups: true,
		}
		urlcollection.EnsureIndex(index)
	} else {
		fmt.Printf("Error occured while creating mongodb connection: %s", err.Error())
	}
	return
}

func (c *MongoConnection) CloseConnection(){
	if c.originalSession != nil {
		c.originalSession.Close()
	}
}

func (c* MongoConnection) getSessionAndCollection()(session *mgo.Session, urlCollection *mgo.Collection, err error){
	if c.originalSession != nil {
		session = c.originalSession.Copy()
		urlCollection = session.DB("LinkShortnerDB").C("UrlCollection")
	} else {
		err = errors.New("No original session found")
	}
	return
}

func (c *MongoConnection) FindshortUrl(longurl string) (sUrl string,err error) {
	result := mongoDocument{}
	session, urlCollection,err := c.getSessionAndCollection()
	if err == nil {
		defer session.Close()
		err = urlCollection.Find(bson.M{"longurl" :longurl}).One(&result)
		if err == nil {
			sUrl = result.ShortUrl
		}
	}
	return
}

func (c *MongoConnection) FindlongUrl(shortUrl string) (lUrl string,err error) {
	result := mongoDocument{}
	session, urlCollection,err := c.getSessionAndCollection()
	if err==nil {
		defer session.Close()
		err = urlCollection.Find(bson.M{"shorturl" :shortUrl}).One(&result)
		if err == nil {
			lUrl = result.LongUrl
		}
	}
	return
}

func (c *MongoConnection) AddUrls(longUrl string, shortUrl string) (err error) {
	session, urlCollection,err := c.getSessionAndCollection()
	if err == nil {
		defer session.Close()
		err = urlCollection.Insert(
			&mongoDocument{
				Id: bson.NewObjectId(),
				ShortUrl: shortUrl,
				LongUrl: longUrl,
			},
		)
		if err != nil {
			if mgo.IsDup(err) {
				err = errors.New("Duplicate name exists for the shorturl")
			}
		}
	}
	return
}




