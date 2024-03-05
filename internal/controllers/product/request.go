package product_controller

import "github.com/gin-gonic/gin"

type GetAllProductsRequest struct {
	Page  int64 `form:"page" validate:"required,gt=0"`
	Limit int64 `form:"limit" validate:"required,gt=0"`
}

type GetProductRequest struct {
	Id string `uri:"id" validate:"required,mongodb"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" bson:"name" validate:"required,min=3"`
	Description *string `json:"description" bson:"description" validate:"omitempty,min=10"`
	Price       float64 `json:"price" bson:"price" validate:"required,gt=0"`
	Source      string  `json:"source" bson:"source" validate:"required,min=3"`
	ImageUrl    *string `json:"imageUrl" bson:"imageUrl" validate:"omitempty,url"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" bson:"name" validate:"min=3"`
	Description *string `json:"description" bson:"description" validate:"omitempty,min=10"`
	Price       float64 `json:"price" bson:"price" validate:"gt=0"`
	Source      string  `json:"source" bson:"source" validate:"min=3"`
	ImageUrl    *string `json:"imageUrl" bson:"imageUrl" validate:"omitempty,url"`
}

func (getAllProductsRequest *GetAllProductsRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindQuery(getAllProductsRequest)
}

func (getProductRequest *GetProductRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindUri(getProductRequest)
}

func (createProductRequest *CreateProductRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(createProductRequest)
}

func (updateProductRequest *UpdateProductRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(updateProductRequest)
}
