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
	db, err := gorm.Open(g.DBMS, g.USER+":"+g.PASS+"@tcp("+g.DBHOST+")/"+g.DBNAME+"?"+g.OPTION)
	if err != nil {
		log.Fatalln(err.Error())
	}

	db.SingularTable(true)

	return db
}

func (g *GormMutex) Insert(model interface{}) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Create(model)
}

func (g *GormMutex) SelectByWhereOneColumn(model interface{}, column string, value string) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Where(column+"=?", value).First(model)
}

func (g *GormMutex) Count(model interface{}, column string, value string) (*gorm.DB, int) {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	var c int
	return db.Find(model).Where(column+"=?", value).Count(&c), c
}

func (g *GormMutex) SelectAll(model interface{}) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Find(model)
}

func (g *GormMutex) Update(model interface{}) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	db.Model(model).Updates(model)

	return db.Save(model)
}

func (g *GormMutex) Delete(model interface{}) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Delete(model)
}

func (g *GormMutex) SelectByWhereIn(model interface{}, column string, list []string) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := g.connect()
	log.Printf("MySQL Connect Success: %s", runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Where(column+" IN (?)", list).Find(model)
}
