package service

import (
	"crobe-ecommerce/app/backend/internal/product/models"
	"crobe-ecommerce/app/backend/internal/product/repository"
)

// --|| Aqui posso colocar o repo do DB e etc ||--\\
type service struct {
	repo repository.Repository
}

// DeleteProduct implements Service.
func (s *service) DeleteProduct(product *models.DeleteRequest) error {
	return s.repo.DeleteProduct(product)
}

// GetProduct implements Service.
func (s *service) GetProduct(id string) (*models.Product, error) {
	return s.repo.FetchProduct(id)
}

// InsertBook implements Service.
func (s *service) InsertProduct(product *models.Product) (*models.Product, error) {
	return s.repo.CreateProduct(product)
}

// --|| Adicionar a interface aqui ||--\\
type Service interface {
	InsertProduct(product *models.Product) (*models.Product, error)
	GetProduct(id string) (*models.Product, error)
	DeleteProduct(product *models.DeleteRequest) error
}

func NewService(r repository.Repository) Service {
	return &service{
		repo: r,
	}
}
