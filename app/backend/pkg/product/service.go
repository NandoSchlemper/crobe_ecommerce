package product

import "crobe-ecommerce/app/backend/pkg/entities"

//--|| Aqui posso colocar o repo do DB e etc ||--\\
type service struct {
	repo Repository
}

// InsertBook implements Service.
func (s *service) InsertBook(product *entities.Product) (*entities.Product, error) {
	return s.repo.CreateProduct(product)
}

//--|| Adicionar a interface aqui ||--\\
type Service interface {
	InsertBook(product *entities.Product) (*entities.Product, error)
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
