package handlers

import (
	dto "housy/dto/result"
	"housy/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handlerFilter struct {
	FilterRepository repositories.FilterRepository
}

func HandlerFilter(FilterRepository repositories.FilterRepository) *handlerFilter {
	return &handlerFilter{FilterRepository}
}

func (h *handlerFilter) MultiFilter(c echo.Context) error {
	params := c.QueryParams()

	houses, err := h.FilterRepository.MultiFilter(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := dto.SuccessResult{Code: http.StatusOK, Data: houses}
	return c.JSON(http.StatusOK, response)
}
