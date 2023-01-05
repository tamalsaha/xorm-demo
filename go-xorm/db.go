package go_xorm

import (
	"fmt"
	"log"
	"xorm.io/core"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	//"xorm.io/xorm"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "postgres"
	NAME     = "xorm-demo"
)

var en *xorm.Engine

func Connect() *xorm.Engine {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, NAME)
	var err error
	en, err = xorm.NewEngine("postgres", dbinfo)
	if err != nil {
		log.Println("engine creation failed", err)
	}

	en.SetMapper(core.GonicMapper{})
	// defer post.db.Close()

	en.ShowSQL(true)
	err = en.Ping()
	if err != nil {
		panic(err)
	}

	syncTables(en)
	log.Println("Successfully connected")
	return en
}

func Close() {
	en.Close()
}

func syncTables(en *xorm.Engine) {
	err := en.Sync(
		new(User),
		new(Contract),
	)
	if err != nil {
		log.Println("creation error", err)
		return
	}
	log.Println("Successfully synced")
}
