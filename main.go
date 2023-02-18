package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

type Test struct {
	gorm.Model
	Name string `gorm:"column:test_name;type:varchar(50);unique;default:unknown;not null"`
}

type Gender struct {
	ID uint
	Name string `gorm:"unique;size(10)"`
}

func (t Test) TableName() string {
	return "my_test";
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc();
	fmt.Printf("%v\n=====================\n", sql);
}

var db *gorm.DB;

func CreateGender(name string) {
	gender := Gender{Name: name};
	tx := db.Create(&gender);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return; 
	}
	fmt.Println(gender); 
}

func GetGenders() {
	genders := []Gender{};
	tx := db.Order("id").Find(&genders);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}
	fmt.Println(genders);
}

func GetGenderByID(id int) {
	gender := Gender{};
	tx := db.First(&gender, id);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}
	fmt.Println(gender);
}

func GetGenderByName(name string) {
	gender := Gender{};
	tx := db.First(&gender, "name=?", name);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}
	fmt.Println(gender);
}

func main() {
	dsn := "root:P@ssw0rd@tcp(0.0.0.0:3306)/goorm_codebangkok?parseTime=true";
	mysql := mysql.Open(dsn);
	var err error;
	db, err = gorm.Open(mysql, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	});
	if err != nil {
		panic(err);
	}
	
	// err = db.AutoMigrate(Gender{}, Test{});
	// if err != nil {
	// 	fmt.Println(err); 
	// }

	GetGenders(); 
	GetGenderByID(2); 
	GetGenderByName("Male");
}
