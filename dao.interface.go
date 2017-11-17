//interface
package main

type AdminDao interface {
	querySingleRowById(id string) AdminDao
	queryRows() AdminDaos
	isTrue(id string) bool
}
type AdminDaos []AdminDao
