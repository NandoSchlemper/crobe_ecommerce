package product

import "crobe-ecommerce/app/backend/pkg/entities"

//--|| Aqui posso colocar o repo do DB e etc ||--\\
type service struct {
	repo Repository
}

// DeleteProduct implements Service.
func (s *service) DeleteProduct(product *entities.DeleteRequest) error {
	return s.repo.DeleteProduct(product)
}

// GetProduct implements Service.
func (s *service) GetProduct(id string) (*entities.Product, error) {
	return s.repo.FetchProduct(id)
}

// InsertBook implements Service.
func (s *service) InsertProduct(product *entities.Product) (*entities.Product, error) {
	return s.repo.CreateProduct(product)
}

//--|| Adicionar a interface aqui ||--\\
type Service interface {
	InsertProduct(product *entities.Product) (*entities.Product, error)
	GetProduct(id string) (*entities.Product, error)
	DeleteProduct(product *entities.DeleteRequest) error
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
