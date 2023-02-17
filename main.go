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

type Gender struct {
	ID uint
	Name string
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc();
	fmt.Printf("%v\n=====================\n", sql);
}

func main() {
	dsn := "root:P@ssw0rd@tcp(0.0.0.0:3306)/goorm_codebangkok?parseTime=true";
	mysql := mysql.Open(dsn);
	db, err := gorm.Open(mysql, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: true,
	});
	if err != nil {
		panic(err);
	}
	
	err = db.Migrator().CreateTable(Gender{});
	if err != nil {
		fmt.Println(err); 
	}

}
