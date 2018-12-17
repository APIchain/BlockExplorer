package module

import (
	db "github.com/BlockExplorer/db"
	mgo "gopkg.in/mgo.v2"
)

func UserCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("users")
}

func BlockCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("Blocks")
}

func AccountCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("Accounts")
}

func TransactionCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("Transactions")
}
