package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/alissonFabricio04/ecommerce/backend/internal/application/usecases"
	"github.com/alissonFabricio04/ecommerce/backend/internal/infra/api/utils"
	"github.com/alissonFabricio04/ecommerce/backend/internal/infra/repositories"
)

type CreateNewCategoryReq struct {
	Name string `json:"name"`
}

type CreateNewCategoryResSucess struct {
	CategoryId string `json:"categoryId"`
}

type CreateNewCategoryResError struct {
	Message string `json:"message"`
}

func CreateNewCategoryController(w http.ResponseWriter, r *http.Request) {
	body, err := utils.BodyReader(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := CreateNewCategoryResError{Message: "error when parsed the body"}
		utils.SendResponse(w, response)
		return
	}

	var category CreateNewCategoryReq
	if err = json.Unmarshal(body, &category); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := CreateNewCategoryResError{Message: "error decoding JSON"}
		utils.SendResponse(w, response)
		return
	}

	usecase := usecases.InstaceNewCreateNewCategoryUseCase(repositories.NewCategoryRepositoryImpl())
	categoryId, err := usecase.Handle(category.Name)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		response := CreateNewCategoryResError{Message: err.Error()}
		utils.SendResponse(w, response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := CreateNewCategoryResSucess{CategoryId: categoryId.ToString()}
	utils.SendResponse(w, response)
}
