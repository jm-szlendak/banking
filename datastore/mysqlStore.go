package datastore

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jm-szlendak/banking/models"
)

type MySqlTransationStore struct {
	dbURI          string
	dbName         string
	collectionName string
	db             *sql.DB
}

func (s *MySqlTransationStore) Open() *MySqlTransationStore {
	db, err := sql.Open("mysql", "kuba:kuba@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		log.Fatal(err)
	}
	s.db = db

	return s
}

func (s *MySqlTransationStore) Close() error {
	return s.db.Close()
}

func (s *MySqlTransationStore) GetAll() ([]models.Transaction, error) {
	return make([]models.Transaction, 0, 0), nil
}

func (s *MySqlTransationStore) Insert(transactions []models.Transaction, replace bool) error {
	return nil
}

func NewMySqlTransationStore(dbURI, database, collection string) *MySqlTransationStore {
	var store = MySqlTransationStore{
		dbURI:          dbURI,
		dbName:         database,
		collectionName: collection,
	}
	return store.Open()
}
