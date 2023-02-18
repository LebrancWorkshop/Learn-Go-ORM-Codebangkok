package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

type Test struct {
	gorm.Model
	Code uint
	Name string `gorm:"column:test_name;type:varchar(50);unique;default:unknown;not null"`
}

type Customer struct {
	ID uint
	Name string
	Gender Gender
	GenderID uint
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

func GetGenderByID(id uint) {
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
	tx := db.Where("name=?", name).Find(&gender);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}
	fmt.Println(gender);
}

func UpdateGender(id uint, newName string) {
	gender := Gender{};
	tx := db.First(&gender, id);
	if tx.Error != nil {
		fmt.Println(tx.Error);
	}
	gender.Name = newName;
	
	tx = db.Save(&gender);
	if tx.Error != nil {
		fmt.Println(tx.Error);
	}
}

func UpdateGenderByName(id uint, newName string) {
	gender := Gender{Name: newName};
	tx := db.Model(&Gender{}).Where("id=@myid", sql.Named("myid", id)).Updates(gender);
	if tx.Error != nil {
		fmt.Println(tx.Error);
	}
	GetGenderByID(id);
}

func DeleteGender(id int) {
	tx := db.Delete(&Gender{}, id);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}
	fmt.Println("Deleted Success");
}

func CreateTest(code uint, name string) {
	test := Test{Code: code, Name: name};
	tx := db.Create(&test); 
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}
	fmt.Println(test);
}

func GetTests() {
	tests := []Test{};
	tx := db.Find(&tests);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}
	fmt.Println(tests);
	
}

func DeleteTestSoft(id uint) {
	tx := db.Delete(&Test{}, id);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}
}

func DeleteTestHard(id uint) {
	tx := db.Unscoped().Delete(&Test{}, id);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}
}

func CreateCustomer(name string, genderID uint) {
	customer := Customer{
		Name: name,
		GenderID: genderID,
	}

	tx := db.Create(&customer);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}
}

func GetCustomers() {
	customers := []Customer{};
	tx := db.Preload(clause.Associations).Find(&customers);
	if tx.Error != nil {
		fmt.Println(tx.Error);
		return;
	}

	for _, customer := range customers {
		fmt.Printf("%v|%v|%v\n", customer.ID, customer.Name, customer.Gender.Name);
	}
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
	
	// err = db.AutoMigrate(Gender{}, Test{}, Customer{});
	// if err != nil {
	// 	fmt.Println(err); 
	// }

	// CreateCustomer("Chin", 1);
	UpdateGenderByName(2, "Pooying"); 
	GetCustomers(); 
}
