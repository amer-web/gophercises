package app

import (
	"github.com/boltdb/bolt"
	"log"
)

var Myapp *App

type App struct {
	Db *bolt.DB
}

func Run() *App {
	if Myapp == nil {
		Myapp = &App{}
	}
	db, err := bolt.Open("amer.db", 0600, nil)
	Myapp.Db = db
	if err != nil {
		log.Fatal(err)
	}
	return Myapp
}

func GetDB() *bolt.DB {
	return Myapp.Db
}
