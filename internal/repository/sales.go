package repository

import (
	"context"
	"fmt"

	"github.com/Fran313/retailBrain/config"
	"github.com/Fran313/retailBrain/internal/model"
	"github.com/jackc/pgx/v5"
)

func InsertSalesBulk(sales []model.Sale) error {
	conn, err := config.DB.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("failed to acquire DB connection: %w", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	for _, s := range sales {
		batch.Queue(
			`INSERT INTO retail.sales
			(store, section, product, product_id, net_sale, net_sale_var_lyc,
			 units, units_ly, units_var_ly, units_lyc, units_var_lyc)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			s.Store, s.Section, s.Product, s.ProductID,
			s.NetSale, s.NetSaleVarLYC,
			s.Units, s.UnitsLY, s.UnitsVarLY,
			s.UnitsLYC, s.UnitsVarLYC,
		)
	}

	br := conn.SendBatch(context.Background(), batch)
	defer br.Close()

	for i := 0; i < len(sales); i++ {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("insert failed at row %d: %w", i, err)
		}
	}

	return nil
}
