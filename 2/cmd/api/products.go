package main
import (
	"fmt"
	"net/http"
	"greenworld.wow/internal/data"
	"greenworld.wow/internal/validator"
)
func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Category string `json:"category"`
		Price int8 `json:"price"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Copy the values from the input struct to a new Product struct.
	product := &data.Product{
		Title: input.Title,
		Category: input.Category,
		Price: input.Price,
	}
	// Initialize a new Validator.
	v := validator.New()
	// Call the ValidateProduct() function and return a response containing the errors if
	// any of the checks fail.
	if data.ValidateProduct(v, product); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
// Add a showProductHandler for the "GET /v1/products/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.

func (app *application) showProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		// Use the new notFoundResponse() helper.
		app.notFoundResponse(w, r)
		return
	}
	product := data.Product{
		ID: id,
		Title: "EcoPaper",
		Category: "Recycled Paper Products",
		Price: 100,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"product": product}, nil)
	if err != nil {
		// Use the new serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
	}
}
