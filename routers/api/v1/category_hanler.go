package v1

import (
	"file-server/pkg/app"
	"file-server/pkg/e"
	"file-server/services/category_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategoryHandler(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, "this is GetCategoryHandler")
}

type AddCategoryForm struct {
	CategoryTitle string `form:"category_title" valid:"Required"`
	CategoryDesc  string `form:"category_desc"`
}

// @Summary Add Category
// @Produce  json
// @Param category_title body string true "CategoryTitle"
// @Param category_desc body string true  "CategoryDesc"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/category [post]
func AddCategoryHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddCategoryForm
	)

	err := c.ShouldBind(&form)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error)
		return
	}

	categoryService := category_service.Category{
		CategoryTitle: form.CategoryTitle,
		CategoryDesc:  form.CategoryDesc,
	}

	err = categoryService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteCategoryHandler(c *gin.Context){
	var (
		appG = app.Gin{C: c}
	)
	categoryNo := c.Param("category_no")

	categoryService := category_service.Category{
		CategoryNo: categoryNo,
	}

	err := categoryService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
