-- Crear schema si no existe
CREATE SCHEMA IF NOT EXISTS retail;

-- Crear tabla sales
CREATE TABLE IF NOT EXISTS retail.sales (
    id SERIAL PRIMARY KEY,
    store TEXT NOT NULL,
    section TEXT NOT NULL,
    product TEXT NOT NULL,
    product_id INTEGER NOT NULL,
    net_sale DOUBLE PRECISION NOT NULL,
    net_sale_var_lyc DOUBLE PRECISION NOT NULL,
    units INTEGER NOT NULL,
    units_ly INTEGER NOT NULL,
    units_var_ly DOUBLE PRECISION NOT NULL,
    units_lyc INTEGER NOT NULL,
    units_var_lyc DOUBLE PRECISION NOT NULL
); 