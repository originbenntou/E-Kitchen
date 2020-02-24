package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"runtime"
	"sync"
)

type GormMutex struct {
	lock    sync.RWMutex
	counter uint64

	DBMS   string
	USER   string
	PASS   string
	DBHOST string
	DBNAME string
	OPTION string
}

func (g *GormMutex) connect() *gorm.DB {
	CONNECT := g.USER + ":" + g.PASS + "@tcp(" + g.DBHOST + ")/" + g.DBNAME + "?" + g.OPTION

	db, err := gorm.Open(g.DBMS, CONNECT)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}

func (g *GormMutex) Insert(i interface{}) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Create(i)
}

func (g *GormMutex) Select(table interface{}, column string, value string) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Find(table).Where(column+"=?", value)
}

func (g *GormMutex) Count(table interface{}, column string, value string) (*gorm.DB, int) {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	var c int
	return db.Find(table).Where(column+"=?", value).Count(&c), c
}

func (g *GormMutex) SelectAll(table interface{}) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Find(table)
}

func (g *GormMutex) Update(table interface{}) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Save(table)
}
