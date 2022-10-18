package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/bxcodec/faker/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// A model is a structure reflecting a database table structure. An instance of a model
// is a single database record. Each model is defined in its own file inside the database/models directory.
// Models are usually just normal Golang structs, basic Go types, or pointers of them.
// "sql.Scanner" and "driver.Valuer" interfaces are also supported.

// Learn more here: https://goyave.dev/guide/basics/database.html#models

var globalDB *gorm.DB

func Conn() *gorm.DB {
	return globalDB
}

func init() {
	// All models should be registered in an "init()" function inside their model file.

	// postgres://vqlcxppersyinr:e9471d873dada65c370c7ec26e2636410a9890f0de257c91c851cd31b81ba820@ec2-63-32-248-14.eu-west-1.compute.amazonaws.com:5432/d7tute07d3u68g
	// database.RegisterDialect("postgres", "postgres://{username}:{password}@{host}:{port}/{name}", postgres.Open)
	// https://github.com/go-gorm/postgres
	log.Println(os.Environ())

	sqlDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Sql DB ", sqlDB)

	//database.RegisterModel(&Person{})

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	migre := gormDB.Migrator()
	migre.AutoMigrate(&Person{})

	globalDB = gormDB
}

type Person struct {
	gorm.Model
	Name    string
	Age     int32
	Address string
	Work    string
}

func PersonGenerator() interface{} {
	person := &Person{}
	person.Name = faker.Name()
	test, _ := faker.RandomInt(18, 72)
	person.Age = int32(test[0])
	person.Address = "45s"
	person.Work = "сварщик"
	//faker.SetGenerateUniqueValues(true)
	//user.Email = faker.Email()
	//faker.SetGenerateUniqueValues(false)
	return person
}

// func ConnectDB() *gorm.DB {
// 	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=musicstore password=postgres sslmode=disable")
// 	if err != nil {
// 		panic("Не удалось подключиться к базе данных")
// 	}
// 	db.AutoMigrate(&Track{})

//		return db
//	}
//
// User represents a user.
type User struct {
	gorm.Model
	Name  string `gorm:"type:char(100)"`
	Email string `gorm:"type:char(100);uniqueIndex"`
}

// You may need to test features interacting with your database.
// Goyave provides a handy way to generate and save records in your database: factories.
// Factories need a generator function. These functions generate a single random record.
//
// "database.Generator" is an alias for "func() interface{}"
//
// Learn more here: https://goyave.dev/guide/advanced/testing.html#database-testing

// UserGenerator generator function for the User model.
// Generate users using the following:
//
//	database.NewFactory(model.UserGenerator).Generate(5)
func UserGenerator() interface{} {
	user := &User{}
	user.Name = faker.Name()

	faker.SetGenerateUniqueValues(true)
	user.Email = faker.Email()
	faker.SetGenerateUniqueValues(false)
	return user
}
