package person

import (
	"net/http"
	"strconv"

	"github.com/DrStarland/probagoyave/database/model"
	"goyave.dev/goyave/v4"
	"goyave.dev/goyave/v4/database"
)

func Hohoho(response *goyave.Response, request *goyave.Request) {
	response.String(http.StatusOK, "ohohoh")
}

// Method	URI	Handler name	Description
// GET	/product	Index()	Get the products list
// POST	/product	Store()	Create a product
// GET	/product/{id}	Show()	Show a product
// PUT or PATCH	/product/{id}	Update()	Update a product
// DELETE	/product/{id}	Destroy()	Delete a product

// SayHi is a controller handler writing "Hi!" as a response.
//
// The Response object is used to write your response.
// https://goyave.dev/guide/basics/responses.html
//
// The Request object contains all the information about the incoming request, including it's parsed body,
// query params and route parameters.
// https://goyave.dev/guide/basics/requests.html
func AllPersons(response *goyave.Response, request *goyave.Request) {
	// temp := dbmod.Person{}
	db := database.GetConnection()
	//db.AutoMigrate(&dbmod.Person{})
	// Create
	// pers := dbmod.PersonGenerator().(*dbmod.Person)
	// db.Create(pers)
	// db.Raw("select * from \"people\" where \"people\".\"id\" == 5").Scan(&temp)
	// Read
	// db.First(&temp, 7) // find product with integer primary key

	//log.Println(db)

	var results []map[string]interface{}
	//var results []dbmod.Person
	//db.Table("people").Find(&results)
	db.Model(&model.Person{}).Find(&results)

	response.JSON(http.StatusOK, results)
}

func GetPersonByID(response *goyave.Response, request *goyave.Request) {
	// temp := dbmod.Person{}
	strID := request.Params["personId"]
	ID, _ := strconv.Atoi(strID)
	db := database.GetConnection()
	//db.AutoMigrate(&dbmod.Person{})
	// Create
	// pers := dbmod.PersonGenerator().(*dbmod.Person)
	// db.Create(pers)
	// db.Raw("select * from \"people\" where \"people\".\"id\" == 5").Scan(&temp)
	// Read
	// db.First(&temp, 7) // find product with integer primary key

	//log.Println(db)

	var results []map[string]interface{}
	//var results []dbmod.Person
	db.Table("people").Find(&results, "id = ?", ID)
	//db.Model(&model.Person{}).Select("*").Where()
	response.JSON(http.StatusOK, results)
}

func CreatePerson(response *goyave.Response, request *goyave.Request) {
	// product := model.Product{
	// 	Name:  request.String("name"),
	// 	Price: request.Numeric("price"),
	// }
	// if err := database.Conn().Create(&product).Error; err != nil {
	// 	response.Error(err)
	// } else {
	// 	response.JSON(http.StatusCreated, map[string]uint{"id": product.ID})
	// }
	person := model.Person{
		Name:    request.String("name"),
		Age:     int32(request.Numeric("age")),
		Address: request.String("address"),
		Work:    request.String("work"),
	}
	if err := database.Conn().Create(&person).Error; err != nil {
		response.Error(err)
	} else {
		response.JSON(http.StatusCreated,
			map[string]uint{
				"id": person.ID,
			})
	}
}

// func Echo(response *goyave.Response, request *goyave.Request) {
// 	response.String(http.StatusOK, request.String("text"))
// }
