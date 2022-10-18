package person

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DrStarland/probagoyave/database/model"
	"goyave.dev/goyave/v4"
)

func Hohoho(response *goyave.Response, request *goyave.Request) {
	response.String(http.StatusOK, "ohohoh")
}

// PUT or PATCH	/product/{id}	Update()	Update a product
// DELETE	/product/{id}	Destroy()	Delete a product

// Получить информацию по всем людям
func Index(response *goyave.Response, request *goyave.Request) {
	// db := database.GetConnection()
	db := model.Conn()
	var results []map[string]interface{}
	db.Model(&model.Person{}).Find(&results)
	response.JSON(http.StatusOK, results)
}

// Получить информацию о человеке по его ID
func Show(response *goyave.Response, request *goyave.Request) {
	// strID := request.Params["personId"]
	// ID, _ := strconv.Atoi(strID)
	// db := database.GetConnection()
	// var results []map[string]interface{}
	// db.Table("people").Find(&results, "id = ?", ID)
	// response.JSON(http.StatusOK, results)
	pers := model.Person{}
	db := model.Conn()
	result := db.First(&pers, request.Params["personID"])
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, pers)
	}
}

// Метод создания новой записи о человеке
func Store(response *goyave.Response, request *goyave.Request) {
	log.Println(request.Data)
	db := model.Conn()
	person := model.Person{
		Name:    request.String("name"),
		Age:     int32(request.Integer("age")),
		Address: request.String("address"),
		Work:    request.String("work"),
	}
	if err := db.Create(&person).Error; err != nil {
		response.Error(err)
	} else {
		// response.JSON(http.StatusCreated,
		// 	// map[string]uint{
		// 	// 	"id": person.ID,
		// 	// })
		response.Status(http.StatusCreated)
		response.Header().Add("Location", fmt.Sprintf("/api/v1/persons/%d", person.ID))
	}
}

// Метод обновления информации о человеке
func Update(response *goyave.Response, request *goyave.Request) {
	pers := model.Person{}
	db := model.Conn()
	result := db.Select("id").First(&pers, request.Params["personID"])
	if response.HandleDatabaseError(result) {
		_, ageExist := request.Data["age"]
		_, nameExist := request.Data["name"]
		_, addrExist := request.Data["address"]
		_, workExist := request.Data["work"]

		if !(ageExist || nameExist || addrExist || workExist) {
			response.Status(http.StatusBadRequest)
			return
		}

		if nameExist {
			name := request.String("name")
			if err := db.Model(&pers).Update("name", name).Error; err != nil {
				response.Error(err)
			}
		}

		if ageExist {
			age := int32(request.Integer("age"))
			if err := db.Model(&pers).Update("age", age).Error; err != nil {
				response.Error(err)
			}
		}

		if workExist {
			work := request.String("work")
			if err := db.Model(&pers).Update("work", work).Error; err != nil {
				response.Error(err)
			}
		}

		if addrExist {
			address := request.String("address")
			if err := db.Model(&pers).Update("address", address).Error; err != nil {
				response.Error(err)
			}
		}
	}
	response.JSON(http.StatusOK, pers)
}

// Метод удаления информации о человеке
func Destroy(response *goyave.Response, request *goyave.Request) {
	pers := model.Person{}
	db := model.Conn()
	result := db.Select("id").First(&pers, request.Params["personID"])
	if response.HandleDatabaseError(result) {
		if err := db.Delete(&pers).Error; err != nil {
			response.Error(err)
			log.Println(err.Error())
		}
	}

	// Person for ID was removed
	response.WriteHeader(http.StatusNoContent)
}

// 	// temp := dbmod.Person{}
// 	db := database.GetConnection()
// 	//db.AutoMigrate(&dbmod.Person{})
// 	// Create
// 	// pers := dbmod.PersonGenerator().(*dbmod.Person)
// 	// db.Create(pers)
// 	// db.Raw("select * from \"people\" where \"people\".\"id\" == 5").Scan(&temp)
// 	// Read
// 	// db.First(&temp, 7) // find product with integer primary key

// 	//log.Println(db)

// 	var results []map[string]interface{}
// 	//var results []dbmod.Person
// 	//db.Table("people").Find(&results)
// 	db.Model(&model.Person{}).Find(&results)

// 	response.JSON(http.StatusOK, results)
