package handler

import (
	"agolang/project-3/dto"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) categoryHandler {
	return categoryHandler{
		categoryService: categoryService,
	}
}

func (ch *categoryHandler) CreateCategory(ctx *gin.Context) {
	var requestBody dto.NewCategoryRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	result, err := ch.categoryService.CreateNewCategory(requestBody)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (ch *categoryHandler) GetAllCategories(ctx *gin.Context) {
	categories, err := ch.categoryService.GetAllCategories()
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (ch *categoryHandler) UpdateCategory(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")
	categoryIdInt, err := strconv.ParseInt(categoryId, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Category id should be in unsigned integer")
		ctx.JSON(newError.Status(), newError)
		return
	}

	var requestBody dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessibleEntityError(err.Error())
		ctx.JSON(newError.Status(), newError)
		return
	}

	updatedCategory, err2 := ch.categoryService.UpdateCategory(int(categoryIdInt), &requestBody)
	if err2 != nil {
		ctx.JSON(err2.Status(), err2)
		return
	}

	ctx.JSON(http.StatusOK, updatedCategory)
}

func (c *categoryHandler) DeleteCategory(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")
	categoryIdInt, err := strconv.ParseInt(categoryId, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Category Id should be in unsigned integer")
		ctx.JSON(newError.Status(), newError)
		return
	}

	response, err2 := c.categoryService.DeleteCategory(int(categoryIdInt))
	if err2 != nil {
		ctx.JSON(err2.Status(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
