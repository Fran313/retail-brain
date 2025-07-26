package model

type Sale struct {
	Store         string  // Sucursal
	Section       string  // Sección
	Product       string  // Producto
	ProductID     int     // id_producto
	NetSale       float64 // Venta neta s/IVA
	NetSaleVarLYC float64 // Var Venta Neta s/IVA LYC
	Units         float64 // Unidades
	UnitsLY       int     // Unidades LY
	UnitsVarLY    float64 // Var Unidades LY
	UnitsLYC      int     // Unidades LYC
	UnitsVarLYC   float64 // Var Unidades LYC
}
