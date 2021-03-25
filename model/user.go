package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type TUser struct {
	ID       uint `gorm:"primaryKey;AUTO_INCREMENT" `
	Username string
	Password string
}



func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:CYM1437qi@tcp(rm-2ze08bh0zi3p3z246oo.mysql.rds.aliyuncs.com:3306)/pwd?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true},})
	fmt.Println(db)
	fmt.Println(err)
	//user := TUser{Username: "Jinzhu", Password: "123"}
	//result := db.Create(&user) // 通过数据的指针来创建
	//fmt.Println(user.ID)
	//fmt.Println(result.Error)
	GetUserInfo("123")
}

// 通过密码找到对象
func GetUserInfo(password string) TUser {
	dsn := "root:CYM1437qi@tcp(rm-2ze08bh0zi3p3z246oo.mysql.rds.aliyuncs.com:3306)/pwd?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true},})
	fmt.Println(err)
	user:=TUser{Password: password}
	result := db.First(&user)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Rows())
	fmt.Println(user)
	return user
}