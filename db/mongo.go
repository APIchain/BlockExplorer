package db

import (
	"gopkg.in/mgo.v2"
)

const URL string = "mongodb://127.0.0.1:27017"

var c *mgo.Collection
var session *mgo.Session
var mongo *mgo.Database

func GetMongoDB() *mgo.Database {
	return mongo
}

func InitMongoDB(dbname string) {
	session, _ = mgo.Dial(URL)
	//切换到数据库
	db := session.DB(dbname)
	mongo = db
}

// 切换collection
func ChackCollection(c string) {
	mongo.C(c)
}
