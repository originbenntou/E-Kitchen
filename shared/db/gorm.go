package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"runtime"
)

func connect(funcName string) *gorm.DB {
	// FIXME: ConfigMap,Secretへ移行
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	DBHOST := "e-kitchen-mysql:3306"
	DBNAME := "e_kitchen"
	OPTION := "charset=utf8&parseTime=True"

	CONNECT := USER + ":" + PASS + "@tcp(" + DBHOST + ")/" + DBNAME + "?" + OPTION

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("MySQL Connect Success: %s", funcName)

	return db
}

func Insert(i interface{}) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := connect(runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Create(i)
}

func Select(table interface{}, column string, value string) *gorm.DB {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := connect(runtime.FuncForPC(pt).Name())
	defer db.Close()

	return db.Find(table).Where(column+"=?", value)
}

func Count(table interface{}, column string, value string) (*gorm.DB, int) {
	pt, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("trace failed")
	}

	db := connect(runtime.FuncForPC(pt).Name())
	defer db.Close()

	var c int
	return db.Find(table).Where(column+"=?", value).Count(&c), c
}
