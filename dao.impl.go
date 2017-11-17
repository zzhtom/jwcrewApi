// Impl
package main

func (admin Admin) queryRows() AdminDaos {
	var admins AdminDaos
	sql := "SELECT username,realname,email,telphone,website,sex,age,hobby,province,city,town,birthday,introduction FROM think_admin"
	rows, err := db.Query(sql)
	if err != nil {
		logger.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&admin.UserName, &admin.RealName, &admin.Email, &admin.Telephone, &admin.Website, &admin.Sex, &admin.Age, &admin.Hobby,
			&admin.Province, &admin.City, &admin.Town, &admin.Birthday, &admin.Introduction)
		if err != nil {
			logger.Error(err)
		}
		admins = append(admins, admin)
	}
	err = rows.Err()
	if err != nil {
		logger.Error(err)
	}
	return admins
}
func (admin Admin) querySingleRowById(id string) AdminDao {
	sql := "SELECT username,realname,email,telphone,website,sex,age,hobby,province,city,town,birthday,introduction FROM think_admin where username = ?"
	err := db.QueryRow(sql, id).Scan(&admin.UserName, &admin.RealName, &admin.Email, &admin.Telephone, &admin.Website, &admin.Sex, &admin.Age, &admin.Hobby,
		&admin.Province, &admin.City, &admin.Town, &admin.Birthday, &admin.Introduction)
	if err != nil {
		logger.Error(err)
	}
	return admin
}
