package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.Asseter/internal/model"
	"go.Asseter/internal/repository"
	"go.Asseter/internal/services"
	"gorm.io/gorm"
)

type ExampleController struct {
	db         *gorm.DB
	repository *repository.ExampleRepository
	service    *services.ExampleService
}

//TODO rethink Update and Delete

func NewExampleController(db *gorm.DB, service *services.ExampleService, repository *repository.ExampleRepository) *ExampleController {
	return &ExampleController{
		db:         db,
		service:    service,
		repository: repository,
	}
}

// POST
func (e *ExampleController) Post(ctx *gin.Context) {
	var body model.ExampleRequest
	var example model.Example

	if err := ctx.ShouldBindBodyWithJSON(&body); err != nil {
		e.service.Log("Unable to bind")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	//bind body to model
	example.Email = body.Email
	example.Name = body.Name

	if err := e.repository.Create(e.db, &example); err != nil {
		e.service.Log("Internal Server error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	ctx.JSON(http.StatusOK, model.ExampleResponse{
		ID:    example.ID,
		Name:  example.Name,
		Email: example.Email,
	})
}

// GET
func (e *ExampleController) Get(ctx *gin.Context) {
	var example model.Example
	id := ctx.Param("id")  // -> id = /api/:id

	error := e.repository.Get(e.db, &example, id)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		e.service.Log("Not Found")
		ctx.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	} else if error != nil {
		e.service.Log("Internal Server error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
	} 

	ctx.JSON(http.StatusOK, model.ExampleResponse{
		ID: example.ID, 
		Name: example.Name, 
		Email: example.Email,
	})
}

// LIST
func (e *ExampleController) List(ctx *gin.Context) {
	var examples []model.Example
	var response []model.ExampleResponse

	if err := e.repository.List(e.db, &examples); err != nil {
		e.service.Log("Internal Service Error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	for _, example := range examples {
		response = append(response, model.ExampleResponse{
			ID:    example.ID,
			Name:  example.Name,
			Email: example.Email,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

// UPDATE

//DELETE
