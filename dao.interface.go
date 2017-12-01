//interface
package main

type AdminDao interface {
	querySingleRowById(id string) (AdminDao, error)
	queryRows() ([]interface{}, error)
	isExist(id string) (bool, error)
	add() (bool, error)
	deleteById(id string) (bool, error)
}
type AdminDaos []AdminDao

type DatabaseDao interface {
	getConf(dbt *Database) *Database
}
