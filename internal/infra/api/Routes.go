package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alissonFabricio04/ecommerce/backend/internal/infra/api/controllers"
)

func Routes() {
	http.HandleFunc("/all-category", controllers.GetAllCategoriesController)
	http.HandleFunc("/create-new-category", controllers.CreateNewCategoryController)
	http.HandleFunc("/create-new-product", controllers.CreateNewProductController)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
