package datastore

import (
	"github.com/jm-szlendak/banking/models"
	"gopkg.in/mgo.v2"
)

type MongoTransactionStore struct {
	dbConnection   *mgo.Session
	dbURI          string
	dbName         string
	collectionName string
	db             *mgo.Database
	collection     *mgo.Collection
}

func (s *MongoTransactionStore) Open() *MongoTransactionStore {

	dbConnection, err := mgo.Dial(s.dbURI)
	if err != nil {
		panic(err)
	}
	s.dbConnection = dbConnection
	s.db = dbConnection.DB(s.dbName)
	s.collection = s.db.C(s.collectionName)

	return s
}

func (s *MongoTransactionStore) Close() {
	s.dbConnection.Close()
}
func (s *MongoTransactionStore) GetAll() []models.Transaction {

}

func (s *MongoTransactionStore) Insert(transactions []models.Transaction, replace bool) {

}

func NewMongoTransactionStore(dbURI, database, collection string) *MongoTransactionStore {
	return &MongoTransactionStore{
		dbURI:          dbURI,
		dbName:         database,
		collectionName: collection,
	}
}
