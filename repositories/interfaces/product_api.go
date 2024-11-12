package interfaces

import "github.com/fleimkeipa/case/model"

type ProductAPIRepository interface {
	FindAll(suplierID string) (*model.ProductsResponse, error)
	FindOne(id string) (*model.Product, error)
}