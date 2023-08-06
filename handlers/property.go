package handlers

import (
	"fmt"
	propertiesdto "housy/dto/properties"
	dto "housy/dto/result"
	"housy/models"
	"housy/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

var path_file = "http://localhost:5000/uploads/"

//  http://localhost:5000/uploads/xxx.png

type handlerProperty struct {
	PropertyRepository repositories.PropertyRepository
}

func HandlerProperty(propertyRepository repositories.PropertyRepository) *handlerProperty {
	return &handlerProperty{propertyRepository}
}

func (h *handlerProperty) FindProperties(c echo.Context) error {
	products, err := h.PropertyRepository.FindProperties()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	for i, p := range products {
		products[i].Image = path_file + p.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: products})
}

func (h *handlerProperty) GetProperty(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.PropertyRepository.GetProperty(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	product.Image = path_file + product.Image

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertProductResponse(product)})
}

func (h *handlerProperty) AddProperty(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	price, _ := strconv.Atoi(c.FormValue("price"))
	bedroom, _ := strconv.Atoi(c.FormValue("bedroom"))
	bathroom, _ := strconv.Atoi(c.FormValue("bathroom"))

	request := propertiesdto.PropertyRequest{
		Name_Property: c.FormValue("name_property"),
		City:          c.FormValue("city"),
		Address:       c.FormValue("address"),
		Price:         float64(price),
		TypeRent:      c.FormValue("type_of_rent"),
		Amenities:     datatypes.JSON(c.FormValue("amenities")),
		Bedroom:       bedroom,
		Bathroom:      bathroom,
		Sqf:           c.FormValue("sqf"),
		Description:   c.FormValue("description"),
		Image:         dataFile,
		// User_Id: userId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	property := models.Property{
		Name_Property: request.Name_Property,
		City:          request.City,
		Address:       request.Address,
		Price:         request.Price,
		TypeRent:      request.TypeRent,
		Amenities:     request.Amenities,
		Bedroom:       request.Bedroom,
		Bathroom:      request.Bathroom,
		Sqf:           request.Sqf,
		Description:   request.Description,
		Image:         request.Image,
		// User_Id: request.User_Id,

	}

	property, err = h.PropertyRepository.AddProperty(property)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	property, _ = h.PropertyRepository.GetProperty(property.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK, Data: convertProductResponse(property)})
}

func convertProductResponse(u models.Property) propertiesdto.PropertyResponse {
	return propertiesdto.PropertyResponse{
		Name_Property: u.Name_Property,
		City:          u.City,
		Address:       u.Address,
		Price:         u.Price,
		TypeRent:      u.TypeRent,
		Amenities:     u.Amenities,
		Bedroom:       u.Bedroom,
		Bathroom:      u.Bathroom,
		Sqf:           u.Sqf,
		Description:   u.Description,
		Image:         u.Image,
	}
}
