package service

import (
	"context"
	"errors"
	"your_project/models"
	"your_project/repository"
)

type CustomerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

// Create a new customer
func (s *CustomerService) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	return s.repo.Create(ctx, customer)
}

// Get a customer by ID
func (s *CustomerService) GetCustomer(ctx context.Context, id uint) (*models.Customer, error) {
	customer, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

// Update a customer's details
func (s *CustomerService) UpdateCustomer(ctx context.Context, customer *models.Customer) error {
	return s.repo.Update(ctx, customer)
}

// Delete a customer by ID
func (s *CustomerService) DeleteCustomer(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// Get all customers
func (s *CustomerService) GetAllCustomers(ctx context.Context) ([]*models.Customer, error) {
	return s.repo.GetAll(ctx)
}
