package excel

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Fran313/retailBrain/internal/model"
	"github.com/xuri/excelize/v2"
)

func cleanNumberFormat(value string) string {
	// Elimina espacios al inicio y final
	value = strings.TrimSpace(value)

	// Si está vacío después de trim, retorna "0"
	if value == "" {
		return "0"
	}

	// Elimina el símbolo % si existe
	value = strings.TrimSuffix(value, "%")
	value = strings.TrimSpace(value) // Elimina espacios que quedaron

	// Si queda vacío después de quitar %, retorna "0"
	if value == "" {
		return "0"
	}

	// Elimina separadores de miles (comas)
	value = strings.ReplaceAll(value, ",", "")

	return value
}

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
		if i < 2 {
			continue // Skip first 2 rows (empty + headers)
		}
		if len(row) < 11 {
			continue // Incomplete row
		}

		// Manejo de errores con valores por defecto
		productID, err := strconv.Atoi(strings.TrimSpace(row[3]))
		if err != nil {
			productID = 0
		}

		netSale, err := strconv.ParseFloat(cleanNumberFormat(row[4]), 64)
		if err != nil {
			netSale = 0.0
		}

		netSaleVarLYC, err := strconv.ParseFloat(cleanNumberFormat(row[5]), 64)
		if err != nil {
			netSaleVarLYC = 0.0
		}

		units, err := strconv.ParseFloat(cleanNumberFormat(row[6]), 64)
		if err != nil {
			units = 0.0
		}

		unitsLY, err := strconv.Atoi(strings.TrimSpace(row[7]))
		if err != nil {
			unitsLY = 0
		}

		unitsVarLY, err := strconv.ParseFloat(cleanNumberFormat(row[8]), 64)
		if err != nil {
			unitsVarLY = 0.0
		}

		unitsLYC, err := strconv.Atoi(strings.TrimSpace(row[9]))
		if err != nil {
			unitsLYC = 0
		}

		unitsVarLYC, err := strconv.ParseFloat(cleanNumberFormat(row[10]), 64)
		if err != nil {
			unitsVarLYC = 0.0
		}

		sale := model.Sale{
			Store:         strings.ToUpper(strings.TrimSpace(row[0])),
			Section:       strings.ToUpper(strings.TrimSpace(row[1])),
			Product:       strings.ToUpper(strings.TrimSpace(row[2])),
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
