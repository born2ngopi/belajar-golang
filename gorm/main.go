package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Saved struct {
	ID        uint
	Name      string
	Email     string
	ActivedAt sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	gorm.Model
	CreditCard CreditCard `gorm:"foreignKey:UserName"`
	// use UserName as foreign key
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "pass"
	dbname   = "dbGateway"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	gormDb.AutoMigrate(&User{}, &CreditCard{})
}
