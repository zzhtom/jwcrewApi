//interface
package main

type Dao interface {
	add() (bool, error)
	update() (bool, error)
	isExist(id string) (bool, error)
	queryRows() ([]interface{}, error)
	deleteById(id string) (bool, error)
	querySingleRowById(id string) (Dao, error)
}
type Daos []Dao

type DatabaseDao interface {
	getConf(dbt *Database) *Database
}
