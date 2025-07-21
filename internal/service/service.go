package service

import (
	"context"
	"fmt"

	"github.com/Fran313/retailBrain/internal/database"
	"github.com/Fran313/retailBrain/internal/model"
	"github.com/Fran313/retailBrain/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SalesService handles business logic for sales operations
type SalesService struct {
	db *pgxpool.Pool
}

// NewSalesService creates a new sales service
func NewSalesService() *SalesService {
	return &SalesService{
		db: database.DB,
	}
}

// InsertSalesBulk inserts multiple sales records
func (s *SalesService) InsertSalesBulk(ctx context.Context, sales []model.Sale) error {
	// Validate sales data
	for i, sale := range sales {
		if err := s.validateSale(&sale); err != nil {
			return fmt.Errorf("invalid sale data at index %d: %w", i, err)
		}
	}

	// Insert sales in database
	if err := repository.InsertSalesBulk(sales); err != nil {
		return fmt.Errorf("failed to insert sales: %w", err)
	}

	return nil
}

// validateSale validates sale data
func (s *SalesService) validateSale(sale *model.Sale) error {
	if sale == nil {
		return fmt.Errorf("sale cannot be nil")
	}

	if sale.NetSale < 0 {
		return fmt.Errorf("net sale cannot be negative")
	}

	if sale.Product == "" {
		return fmt.Errorf("product cannot be empty")
	}

	if sale.Store == "" {
		return fmt.Errorf("store cannot be empty")
	}

	if sale.Section == "" {
		return fmt.Errorf("section cannot be empty")
	}

	return nil
}
