package db

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// real project use the different host
// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
// const (
// 	host     = "postgresdb"
// 	port     = 5432
// 	user     = "parkadmin"
// 	password = "secret"
// 	dbname   = "parking"
// )

func DB() *gorm.DB {
	env := LoadENV()
	fmt.Println(env.POSTGRES_PORT)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", env.POSTGRES_HOST, StringToInt(env.POSTGRES_PORT), env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB)

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func Migrate() {
	// db := DB()
	// db.AutoMigrate(&UserDB{})
}

func StringToInt(s string) int {
	// string to int
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
