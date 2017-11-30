// Impl
package main

import (
	"errors"
	"io/ioutil"

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
func (admin Admin) queryRows() (AdminDaos, error) {
	sql := "SELECT username,realname,email,telphone,website,sex,age,hobby,province,city,town,birthday,introduction FROM think_admin"

	return find(sql, admin)
}
func (admin Admin) querySingleRowById(id string) (AdminDao, error) {
	sql := "SELECT username,realname,email,telphone,website,sex,age,hobby,province,city,town,birthday,introduction FROM think_admin where username = ?"

	return findOne(sql, id, admin)
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

func findOne(sql string, id string, data interface{}) (AdminDao, error) {
	slice := make([]interface{}, 1)
	switch data.(type) {
	case Admin:
		admin := data.(Admin)
		err := db.QueryRow(sql, id).Scan(&admin.UserName, &admin.RealName, &admin.Email, &admin.Telephone, &admin.Website, &admin.Sex, &admin.Age, &admin.Hobby,
			&admin.Province, &admin.City, &admin.Town, &admin.Birthday, &admin.Introduction)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return admin, nil
	}

	return nil, errors.New("findOne: data type is undefine")
}

func find(sql string, data interface{}) (AdminDaos, error) {
	var admins AdminDaos
	rows, err := db.Query(sql)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()
	switch data.(type) {
	case Admin:
		admin := data.(Admin)
		for rows.Next() {
			err := rows.Scan(&admin.UserName, &admin.RealName, &admin.Email, &admin.Telephone, &admin.Website, &admin.Sex, &admin.Age, &admin.Hobby,
				&admin.Province, &admin.City, &admin.Town, &admin.Birthday, &admin.Introduction)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			admins = append(admins, admin)
		}
	}
	err = rows.Err()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return admins, nil
}
