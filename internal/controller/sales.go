package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Fran313/retailBrain/internal/excel"
	"github.com/Fran313/retailBrain/internal/repository"
	"github.com/gin-gonic/gin"
)

func UploadExcelHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Guardar archivo temporalmente
	// TODO: Guardar en S3
	savePath := filepath.Join(os.TempDir(), file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Leer el Excel
	sales, err := excel.ReadSalesFromExcel(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Excel", "detail": err.Error()})
		return
	}

	// Insert into database
	if err := repository.InsertSalesBulk(sales); err != nil {
		fmt.Printf("❌ INSERT FAILED: %v\n", err) // 👈 Agregá esto
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to insert sales",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales uploaded successfully", "rows": len(sales)})
}
