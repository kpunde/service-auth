package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"serviceAuth/backend/entity/rest"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
	"serviceAuth/backend/service"
	"serviceAuth/backend/utility"
)

var (
	_personService = service.NewPersonService()
)

type PersonController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
}

type personController struct {
}

func (i *personController) FindAll(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	ctx.JSON(200, _personService.FindAll(reqContext))
}

func (i *personController) FindById(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	id := ctx.Param("id")
	ppl, err := _personService.FindById(reqContext, id)
	if !utility.CheckErrorResponse(ctx, 400, err) {
		ctx.JSON(200, ppl)
	}
}

func (i *personController) Save(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	var personRequest rest.NewPersonRequest
	err := ctx.ShouldBindJSON(&personRequest)
	if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
		log.Debug(personRequest)
		person, err := _personService.Save(reqContext, &personRequest)
		if !utility.CheckErrorResponse(ctx, 400, err) {
			ctx.JSON(201, person)
		}
	}
}

func (i *personController) Update(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	var person sql.Person
	err := ctx.ShouldBindJSON(&person)
	if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
		id, err := uuid.Parse(ctx.Param("id"))
		utility.CheckError(err)
		person.Id = id
		if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
			person, err := _personService.Update(reqContext, &person)

			if !utility.CheckErrorResponse(ctx, 400, err) {
				ctx.JSON(http.StatusAccepted, person)
			}
		}
	}
}

func (i *personController) Delete(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	id := ctx.Param("id")
	person, err := _personService.FindById(reqContext, id)
	log.Debug(person)
	if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
		err := _personService.Delete(reqContext, person)
		if !utility.CheckErrorResponse(ctx, 400, err) {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}

func (i *personController) UpdatePassword(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)

	var password sql.PasswordInfo
	err := ctx.ShouldBindJSON(&password)
	if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
		id := ctx.Param("id")
		err := _personService.UpdatePassword(reqContext, id, &password)

		if !utility.CheckErrorResponse(ctx, 400, err) {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}

func NewPersonController() PersonController {
	return &personController{}
}
