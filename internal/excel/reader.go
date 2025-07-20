package excel

import (
	"fmt"
	"strconv"

	"github.com/Fran313/retailBrain/internal/model"
	"github.com/xuri/excelize/v2"
)

func ReadSalesFromExcel(path string) ([]model.Sale, error) {
	file, err := excelize.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file: %w", err)
	}

	sheetName := file.GetSheetName(0)
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read rows: %w", err)
	}

	var sales []model.Sale

	for i, row := range rows {
		if i == 0 {
			continue // Skip header
		}
		if len(row) < 11 {
			continue // Incomplete row
		}

		productID, _ := strconv.Atoi(row[3])
		netSale, _ := strconv.ParseFloat(row[4], 64)
		netSaleVarLYC, _ := strconv.ParseFloat(row[5], 64)
		units, _ := strconv.Atoi(row[6])
		unitsLY, _ := strconv.Atoi(row[7])
		unitsVarLY, _ := strconv.ParseFloat(row[8], 64)
		unitsLYC, _ := strconv.Atoi(row[9])
		unitsVarLYC, _ := strconv.ParseFloat(row[10], 64)

		sale := model.Sale{
			Store:         row[0],
			Section:       row[1],
			Product:       row[2],
			ProductID:     productID,
			NetSale:       netSale,
			NetSaleVarLYC: netSaleVarLYC,
			Units:         units,
			UnitsLY:       unitsLY,
			UnitsVarLY:    unitsVarLY,
			UnitsLYC:      unitsLYC,
			UnitsVarLYC:   unitsVarLYC,
		}
		sales = append(sales, sale)
	}

	return sales, nil
}
