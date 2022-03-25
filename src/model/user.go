package model

import "github.com/jinzhu/gorm"

// 数据表定义
type User struct {
	gorm.Model
	Name      string `gorm: "type: varchar(20); not null"`
	Telephone string `gorm: "type: varchar(100); not null; unique"`
	Password  string `gorm: "size: 255; not null"`
}
