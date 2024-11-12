package repositories

import (
	"net/http"

	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/pkg"
)

type ProductAPIRepository struct {
	Client pkg.Client
}

func NewProductAPIRepository(client pkg.Client) *ProductAPIRepository {
	return &ProductAPIRepository{
		Client: client,
	}
}

func (rc *ProductAPIRepository) FindAll(suplierID string) (*model.ProductsResponse, error) {
	request := model.Request{
		Method:  http.MethodGet,
		Paths:   []string{"suppliers", suplierID, "products"},
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	var products model.ProductsResponse
	if err := rc.Client.Do(request, &products); err != nil {
		return nil, err
	}

	return &products, nil
}

func (rc *ProductAPIRepository) FindOne(id string) (*model.Product, error) {
	return nil, nil
}