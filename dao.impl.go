// Impl
package main

import (
	_ "errors"
	"io/ioutil"
	_ "reflect"

	"gopkg.in/yaml.v2"
)

//for database
func (dbt Database) getConf() Database {

	mysql, err := ioutil.ReadFile("conf.yml")
	if err != nil {
		logger.Error("read yaml file err # ", err)
	}
	err = yaml.Unmarshal(mysql, &dbt)
	if err != nil {
		logger.Error("Unmarshal: %v", err)
	}

	return dbt
}

//for admin
func (admin Admin) queryRows() ([]interface{}, error) {
	sql := "SELECT username,realname,email,telphone,website,sex,age,hobby,province,city,town,birthday,introduction FROM think_admin"
	slice := []interface{}{&admin.UserName, &admin.RealName, &admin.Email, &admin.Telephone, &admin.Website, &admin.Sex, &admin.Age, &admin.Hobby,
		&admin.Province, &admin.City, &admin.Town, &admin.Birthday, &admin.Introduction}

	return find(sql, &admin, slice)
}
func (admin Admin) querySingleRowById(id string) (AdminDao, error) {
	sql := "SELECT username,realname,email,telphone,website,sex,age,hobby,province,city,town,birthday,introduction FROM think_admin where username = ?"
	slice := []interface{}{&admin.UserName, &admin.RealName, &admin.Email, &admin.Telephone, &admin.Website, &admin.Sex, &admin.Age, &admin.Hobby,
		&admin.Province, &admin.City, &admin.Town, &admin.Birthday, &admin.Introduction}
	err := findOne(sql, id, slice)
	return admin, err
}

func (admin Admin) isExist(id string) (bool, error) {
	sql := "SELECT username FROM think_admin where username = ?"

	return exist(sql, id)
}

func (admin Admin) add() (bool, error) {
	sql := "INSERT INTO think_admin (username, realname, email, telphone, website, sex, age, hobby, province, city, town, birthday, introduction) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"

	return tran(sql, admin)
}

func (admin Admin) deleteById(id string) (bool, error) {
	sql := "delete from think_admin where username = ?"

	return tran(sql, id)
}

/**
* from: zhangxioaheng
* data: 2017-11-30
* description:share function
 */

func tran(sql string, data interface{}) (bool, error) {
	tx, err := db.Begin()
	if err != nil {
		logger.Error(err)
		return false, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(sql)
	if err != nil {
		logger.Error(err)
		return false, err
	}
	defer stmt.Close() // danger!
	switch data.(type) {
	case string:
		_, err = stmt.Exec(data)
		if err != nil {
			logger.Error(err)
			return false, err
		}
	case Admin:
		admin := data.(Admin)
		_, err = stmt.Exec(admin.UserName, admin.RealName, admin.Email, admin.Telephone, admin.Website, admin.Sex, admin.Age, admin.Hobby, admin.Province, admin.City, admin.Town, admin.Birthday, admin.Introduction)
		if err != nil {
			logger.Error(err)
			return false, err
		}

	}

	err = tx.Commit()
	if err != nil {
		logger.Error(err)
		return false, err
	}
	// stmt.Close() runs here!
	return true, nil
}

func exist(sql string, id string) (bool, error) {
	temp := ""
	err := db.QueryRow(sql, id).Scan(&temp)
	if err != nil {
		logger.Error(err)
		return false, err
	}

	return true, nil
}

func findOne(sql string, id string, slice []interface{}) error {

	err := db.QueryRow(sql, id).Scan(slice...)
	if err != nil {
		logger.Error(err)
		return err
	}

	return err
}

func find(sql string, admin interface{}, slice []interface{}) ([]interface{}, error) {
	var datas []interface{}
	rows, err := db.Query(sql)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()
	//	switch data.(type) {
	//	case Admin:
	//		admin := data.(Admin)
	//		slice := []interface{}{&admin.UserName, &admin.RealName, &admin.Email, &admin.Telephone, &admin.Website, &admin.Sex, &admin.Age, &admin.Hobby,
	//			&admin.Province, &admin.City, &admin.Town, &admin.Birthday, &admin.Introduction}
	for rows.Next() {
		err := rows.Scan(slice...)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		datas = append(datas, admin)
	}
	//	}
	err = rows.Err()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return datas, nil
}
