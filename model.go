// Model
package main

//for database
type Database struct {
	Dbtype   string `yaml: "dbtype"`
	Username string `yaml: "username"`
	Password string `yaml: "password"`
	Dbname   string `yaml: "dbname"`
	Host     string `yaml: "host"`
	Port     string `yaml: "port"`
	Charset  string `yaml: "charset"`
}

//for table admin
type Admin struct {
	UserName     string `json:"username"`
	RealName     string `json:"realname"`
	Email        string `json:"email"`
	Telephone    string `json:"telphone"`
	Website      string `json:"website"`
	Sex          string `json:"sex"`
	Age          int    `json:"age"`
	Hobby        string `json:"hobby"`
	Province     string `json:"province"`
	City         string `json:"city"`
	Town         string `json:"town"`
	Birthday     string `json:"birthday"`
	Introduction string `json:"introduction"`
}
type Admins []Admin
