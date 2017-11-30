// Dao
package main

import (
	"bytes"
	"database/sql"
	_ "fmt"
	_ "os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/op/go-logging"
)

var lock *sync.Mutex = &sync.Mutex{}
var logger = logging.MustGetLogger("chjw.api")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)
var db *sql.DB = nil

//var once sync.Once

func createInstance(database Database) *sql.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			var err error
			var buf bytes.Buffer
			buf.WriteString(database.Username)
			buf.WriteString(":")
			buf.WriteString(database.Password)
			buf.WriteString("@tcp(")
			buf.WriteString(database.Host)
			buf.WriteString(":")
			buf.WriteString(database.Port)
			buf.WriteString(")/")
			buf.WriteString(database.Dbname)
			buf.WriteString("?charset=")
			buf.WriteString(database.Charset)
			//	once.Do(func() {})
			db, err = sql.Open(database.Dbtype, buf.String())
			if err != nil {
				logger.Error(err)
				return nil
			}
			err = db.Ping()
			if err != nil {
				logger.Error(err)
				return nil
			}
		}
	}

	return db
}
func closeDatabase() {
	if db != nil {
		db.Close()
	}
}
