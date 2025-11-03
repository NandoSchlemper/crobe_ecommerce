package product

import "crobe-ecommerce/app/backend/pkg/entities"

//--|| Aqui posso colocar o repo do DB e etc ||--\\
type service struct {
	repo Repository
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
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
