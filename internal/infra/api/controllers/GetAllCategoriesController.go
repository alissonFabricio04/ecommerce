package controllers

import (
	"net/http"

	"github.com/alissonFabricio04/ecommerce/backend/internal/application/query"
	"github.com/alissonFabricio04/ecommerce/backend/internal/infra/api/utils"
	"github.com/alissonFabricio04/ecommerce/backend/internal/infra/repositories"
)

type GetAllCategoriesRes struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetAllCategoriesController(w http.ResponseWriter, _ *http.Request) {
	query := query.InstaceNewGetAllCategoriesQuery(repositories.NewCategoryRepositoryImpl())
	categories, err := query.Handle()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		response := utils.Response{Message: err.Error()}
		utils.SendResponse(w, response)
		return
	}
	var categoryList []GetAllCategoriesRes
	for _, category := range categories {
		categoryList = append(categoryList, GetAllCategoriesRes{
			Id:   category.Id.ToString(),
			Name: category.Name.Value,
		})
	}
	w.WriteHeader(http.StatusOK)
	response := categoryList
	utils.SendResponse(w, response)
}
