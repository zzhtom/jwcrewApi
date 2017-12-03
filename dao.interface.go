//interface
package main

type Dao interface {
	querySingleRowById(id string) (Dao, error)
	queryRows() ([]interface{}, error)
	isExist(id string) (bool, error)
	add() (bool, error)
	deleteById(id string) (bool, error)
}
type Daos []Dao

type DatabaseDao interface {
	getConf(dbt *Database) *Database
}
