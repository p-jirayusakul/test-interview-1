package http

import (
	"my_project/internal/delivery/http/request"
	"my_project/internal/delivery/http/response"

	"my_project/internal/domain"
	"my_project/internal/usecase"
	"my_project/pkg/error"
	orgresponse "my_project/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductsHandler struct {
	useCase *usecase.ProductUseCase
}

func NewProductHandler(useCase *usecase.ProductUseCase) *ProductsHandler {
	return &ProductsHandler{useCase: useCase}
}

// Create godoc
// @Summary Create a new product
// @Schemes
// @Description Create a new product with name, description, sale price and price
// @Tags product
// @Accept json
// @Produce json
// @Param request body request.CreateProductRequest true "Product information"
// @Success 200 {object} orgresponse.Response{data=response.ProductResponse}
// @Failure 400 {object} orgresponse.Response
// @Failure 500 {object} orgresponse.Response
// @Router /product [post]
func (h *ProductsHandler) Create(c *gin.Context) {

	var req request.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errReq := error.New(error.CodeInvalidInput, err.Error())
		c.JSON(error.HTTPStatus(errReq), orgresponse.ErrorResponse(error.GetErrorCode(errReq)))
		return
	}

	result, err := h.useCase.Create(c.Request.Context(), domain.Product{
		Name:        &req.Name,
		Description: req.Description,
		SalePrice:   &req.SalePrice,
		Price:       &req.Price,
	})
	if err != nil {
		c.JSON(error.HTTPStatus(err), orgresponse.ErrorResponse(error.GetErrorCode(err)))
		return
	}

	c.JSON(http.StatusCreated, makeSuccessProductResponse(result))
}

// Update godoc
// @Summary Update a product
// @Schemes
// @Description Update an existing product with name, description, sale price and price
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body request.UpdateProductRequest true "Product information"
// @Success 200 {object} orgresponse.Response{data=response.ProductResponse}
// @Failure 400 {object} orgresponse.Response
// @Failure 500 {object} orgresponse.Response
// @Router /product/{id} [put]
func (h *ProductsHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errReq := error.New(error.CodeInvalidInput, err.Error())
		c.JSON(error.HTTPStatus(errReq), orgresponse.ErrorResponse(error.GetErrorCode(errReq)))
		return
	}

	var req request.UpdateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errReq := error.New(error.CodeInvalidInput, err.Error())
		c.JSON(error.HTTPStatus(errReq), orgresponse.ErrorResponse(error.GetErrorCode(errReq)))
		return
	}

	result, err := h.useCase.Update(c.Request.Context(), domain.Product{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		SalePrice:   req.SalePrice,
		Price:       req.Price,
	})
	if err != nil {
		c.JSON(error.HTTPStatus(err), orgresponse.ErrorResponse(error.GetErrorCode(err)))
		return
	}

	c.JSON(http.StatusOK, makeSuccessProductResponse(result))
}

// Get godoc
// @Summary Get a product
// @Schemes
// @Description Get a product by ID
// @Tags product
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} orgresponse.Response{data=response.ProductResponse}
// @Failure 400 {object} orgresponse.Response
// @Failure 500 {object} orgresponse.Response
// @Router /product/{id} [get]
func (h *ProductsHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errReq := error.New(error.CodeInvalidInput, err.Error())
		c.JSON(error.HTTPStatus(errReq), orgresponse.ErrorResponse(error.GetErrorCode(errReq)))
		return
	}

	result, err := h.useCase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(error.HTTPStatus(err), orgresponse.ErrorResponse(error.GetErrorCode(err)))
		return
	}

	c.JSON(http.StatusOK, makeSuccessProductResponse(result))
}

func makeSuccessProductResponse(v *domain.Product) orgresponse.Response {

	if v == nil {
		return orgresponse.Response{}
	}

	return orgresponse.Response{
		SuccessFul: true,
		Data: &response.ProductResponse{
			Id:          v.Id,
			Name:        *v.Name,
			Description: v.Description,
			SalePrice:   *v.SalePrice,
			Price:       *v.Price,
		},
	}
}
