package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	mysql.Open();
	gorm.Open();

}