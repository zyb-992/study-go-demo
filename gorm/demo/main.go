package main

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	dsn = "user:password@"
)

func main() {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	ctx := context.Background()
	db = db.WithContext(ctx)
}
