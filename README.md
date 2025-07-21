# Retail Brain

Sistema de análisis de datos de ventas minoristas con capacidades de búsqueda semántica y análisis de tendencias.

## Estructura del Proyecto

```
retail-brain/
├── cmd/
│   └── retail-brain/
│       └── main.go           # Punto de entrada de la aplicación
├── internal/                 # Código privado de la aplicación
│   ├── config/
│   │   └── config.go         # Lógica de configuración
│   ├── database/
│   │   └── db.go             # Conexión y operaciones de base de datos
│   ├── service/
│   │   └── service.go        # Lógica de negocio
│   ├── controller/
│   │   └── sales.go          # Controladores HTTP
│   ├── repository/
│   │   └── sales.go          # Acceso a datos
│   ├── model/
│   │   └── sale.go           # Modelos de datos
│   ├── embeddings/
│   │   └── openai.go         # Integración con OpenAI
│   ├── excel/
│   │   └── reader.go         # Lectura de archivos Excel
│   └── vectorstore/
│       └── qdrant.go         # Integración con Qdrant
├── pkg/                      # Paquetes públicos reutilizables
│   └── common/
│       └── utils.go          # Utilidades comunes
├── api/                      # Definiciones de API
│   └── routes.go             # Rutas de la API
├── web/                      # Archivos web (templates, archivos estáticos)
├── scripts/                  # Scripts auxiliares
│   └── docker-compose.yml    # Configuración de Docker
├── tests/                    # Pruebas de integración
├── go.mod                    # Archivo de módulo Go
├── go.sum                    # Checksums de dependencias
└── README.md                 # Documentación del proyecto
```

## Características

- **Análisis de Ventas**: Procesamiento y análisis de datos de ventas minoristas
- **Búsqueda Semántica**: Capacidades de búsqueda avanzada usando embeddings
- **Base de Datos Vectorial**: Integración con Qdrant para búsquedas semánticas
- **API REST**: Interfaz HTTP para acceder a los datos y funcionalidades
- **Procesamiento de Excel**: Lectura y procesamiento de archivos Excel

## Requisitos

- Go 1.21 o superior
- PostgreSQL
- Qdrant Vector Database
- Docker (opcional)

## Configuración

### Variables de Entorno

```bash
# Base de datos
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin123
DB_NAME=retailBrain

# Servidor
SERVER_PORT=8080

# Qdrant
QDRANT_HOST=localhost
QDRANT_PORT=6333
```

### Instalación

1. Clona el repositorio:
```bash
git clone <repository-url>
cd retail-brain
```

2. Instala las dependencias:
```bash
go mod download
```

3. Configura la base de datos:
```bash
# Usando Docker
docker-compose -f scripts/docker-compose.yml up -d
```

4. Ejecuta la aplicación:
```bash
go run cmd/retail-brain/main.go
```

## Uso

### API Endpoints

- `GET /health` - Estado de salud de la aplicación
- `POST /sales/upload` - Subir archivo Excel con datos de ventas
- `GET /sales` - Obtener lista de ventas
- `POST /search` - Búsqueda semántica en los datos

### Ejemplo de Uso

```bash
# Subir archivo Excel
curl -X POST -F "file=@test-retail-brain.xlsx" http://localhost:8080/sales/upload

# Buscar ventas
curl -X POST -H "Content-Type: application/json" \
  -d '{"query": "ventas de productos electrónicos"}' \
  http://localhost:8080/search
```

## Desarrollo

### Estructura de Código

El proyecto sigue las mejores prácticas de Go:

- **cmd/**: Contiene los puntos de entrada de la aplicación
- **internal/**: Código privado específico de la aplicación
- **pkg/**: Paquetes públicos que pueden ser reutilizados
- **api/**: Definiciones y rutas de la API
- **scripts/**: Scripts de construcción y despliegue

### Convenciones

- Usar `context.Context` para cancelación y timeouts
- Manejar errores con `fmt.Errorf` y wrapping
- Documentar funciones públicas
- Usar interfaces para testing

## Testing

```bash
# Ejecutar todas las pruebas
go test ./...

# Ejecutar pruebas con coverage
go test -cover ./...

# Ejecutar pruebas de integración
go test ./tests/...
```

## Despliegue

### Docker

```bash
# Construir imagen
docker build -t retail-brain .

# Ejecutar contenedor
docker run -p 8080:8080 retail-brain
```

### Producción

1. Configurar variables de entorno para producción
2. Usar un reverse proxy (nginx, traefik)
3. Configurar monitoreo y logging
4. Implementar health checks

## Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles. 