package controllers

import (
	"net/http"

	"github.com/alissonFabricio04/ecommerce/backend/internal/application/query"
	"github.com/alissonFabricio04/ecommerce/backend/internal/infra/api/utils"
	"github.com/alissonFabricio04/ecommerce/backend/internal/infra/repositories"
)

type GetAllCategoriesResSucess struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetAllCategoriesResError struct {
	Message string `json:"message"`
}

func GetAllCategoriesController(w http.ResponseWriter, _ *http.Request) {
	query := query.InstaceNewGetAllCategoriesQuery(repositories.NewCategoryRepositoryImpl())
	categories, err := query.Handle()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		response := GetAllCategoriesResError{Message: err.Error()}
		utils.SendResponse(w, response)
		return
	}
	var categoryList []GetAllCategoriesResSucess
	for _, category := range categories {
		categoryList = append(categoryList, GetAllCategoriesResSucess{
			Id:   category.Id.ToString(),
			Name: category.Name.Value,
		})
	}
	w.WriteHeader(http.StatusOK)
	response := categoryList
	utils.SendResponse(w, response)
}
