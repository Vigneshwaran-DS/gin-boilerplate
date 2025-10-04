# Gin Boilerplate

English | [简体中文](./README_CN.md)

A production-ready Go web application boilerplate built with Gin + GORM, featuring best practices and clean architecture.

## ✨ Features

- 🚀 **Complete Project Structure** - Clean layered architecture (Controller, Service, Model, Router)
- 🔐 **JWT Authentication** - Complete user authentication system
- ⚙️ **Multi-Environment Config** - Support for development, production, and more (Viper-based)
- 🗄️ **Database ORM** - GORM with auto-migration support
- 🔒 **Password Encryption** - Bcrypt password hashing
- 📝 **Logging Middleware** - Request logging
- 🌐 **CORS Support** - Cross-Origin Resource Sharing middleware
- 📦 **Unified Response Format** - Standardized API response structure
- 🎨 **Startup Banner** - Spring Boot-style startup banner
- 🧪 **API Test Scripts** - Multiple testing script options

## 📁 Project Structure

```
gin-boilerplate/
├── config/                 # Configuration files
│   ├── banner.txt         # Startup banner
│   ├── config.go          # Configuration loading logic
│   ├── default.yaml       # Default configuration
│   ├── development.yaml   # Development environment config
│   └── production.yaml    # Production environment config
├── controllers/           # Controller layer
│   ├── auth_controller.go # Authentication controller
│   └── user_controller.go # User controller
├── database/              # Database connection
│   └── database.go
├── middleware/            # Middlewares
│   ├── auth.go           # JWT authentication middleware
│   ├── cors.go           # CORS middleware
│   └── logger.go         # Logging middleware
├── models/                # Data models
│   ├── base.go           # Base model
│   └── user.go           # User model
├── router/                # Router layer
│   └── router.go
├── scripts/               # Scripts
│   ├── api-test.http     # HTTP test file
│   ├── api-test.sh       # Bash test script
│   └── init.sql          # Database initialization script
├── services/              # Business logic layer
│   ├── auth_service.go   # Authentication service
│   └── user_service.go   # User service
├── utils/                 # Utilities
│   ├── banner.go         # Banner utility
│   ├── jwt.go            # JWT utility
│   └── response.go       # Response utility
├── .gitignore
├── go.mod
├── main.go               # Application entry point
└── README.md
```

## 🚀 Quick Start

### Option 1: Docker Deployment (Recommended)

#### Requirements

- Docker
- MySQL (external or separate container)

#### Steps

1. **Clone the project**

```bash
git clone <repository-url>
cd gin-boilerplate
```

2. **Configure production settings**

Edit `config/production.yaml` to match your database configuration:

```yaml
database:
  host: "your-mysql-host"
  port: "3306"
  user: "your-db-user"
  password: "your-db-password"
  dbname: "gin_boilerplate_prod"
```

3. **Build Docker image**

```bash
docker build -t gin-boilerplate:latest .
```

4. **Run container**

```bash
docker run -d \
  --name gin-boilerplate \
  -p 8080:8080 \
  -v $(pwd)/config:/root/config \
  gin-boilerplate:latest
```

5. **Check logs**

```bash
docker logs -f gin-boilerplate
```

6. **Stop container**

```bash
docker stop gin-boilerplate
docker rm gin-boilerplate
```

The service runs on `http://localhost:8080` by default.

### Option 2: Local Development

#### Requirements

- Go 1.19+
- MySQL 5.7+

#### Steps

1. **Clone the Project**

```bash
git clone <repository-url>
cd gin-boilerplate
```

2. **Install Dependencies**

```bash
go mod tidy
```

3. **Database Setup**

**Initialize Database**

```bash
mysql -u root -p < scripts/init.sql
```

**Configure Database Connection**

Edit `config/development.yaml`:

```yaml
database:
  host: "localhost"
  port: "3306"
  user: "root"
  password: "your_password"
  dbname: "gin_boilerplate_dev"

jwt:
  secret: "your-secret-key"
  expire_time: 72
```

4. **Run the Application**

**Development Environment**

```bash
go run main.go
# Or specify environment
go run main.go -e development
```

**Production Environment**

```bash
go run main.go -e production
```

The service runs on `http://localhost:8080` by default.

## 🐳 Docker Deployment Details

### Dockerfile Features

- **Multi-stage build**: Minimizes final image size
- **Alpine-based**: Lightweight and secure
- **Production optimized**: CGO disabled for static binary

### Using with Reverse Proxy

This application is designed to run behind a reverse proxy (Nginx, Traefik, etc.). The reverse proxy should handle:

- SSL/TLS termination
- Load balancing
- Rate limiting
- Static file serving (if needed)

Example Nginx configuration:

```nginx
upstream gin_backend {
    server localhost:8080;
}

server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://gin_backend;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### Production Considerations

Before deploying to production:

1. Change the JWT secret in `config/production.yaml`
2. Configure external MySQL database
3. Set up proper logging and monitoring
4. Configure firewall rules
5. Use reverse proxy for SSL/TLS
6. Set up automated backups for database

## 📚 API Documentation

### Authentication

#### User Registration

```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123",
  "full_name": "Test User"
}
```

**Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "Success",
  "data": {
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "full_name": "Test User",
      "created_at": "2024-01-01T00:00:00Z"
    }
  }
}
```

#### User Login

```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

**Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "Success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "full_name": "Test User"
    }
  }
}
```

### User Endpoints (Authentication Required)

All user-related endpoints require a token in the header:

```http
Authorization: Bearer {your_token}
```

#### Get Current User Info

```http
GET /api/v1/me
Authorization: Bearer {token}
```

#### Update Current User Info

```http
PUT /api/v1/me
Authorization: Bearer {token}
Content-Type: application/json

{
  "full_name": "Updated Name",
  "email": "newemail@example.com"
}
```

#### Update Password

```http
PUT /api/v1/me
Authorization: Bearer {token}
Content-Type: application/json

{
  "password": "newpassword123"
}
```

### Health Check

```http
GET /api/v1/health
```

**Response:**

```json
{
  "status": "ok",
  "message": "Service is running"
}
```

## 🧪 API Testing

The project provides multiple API testing scripts:

### 1. HTTP File Testing (Recommended)

Use VS Code REST Client extension or IntelliJ IDEA HTTP Client:

```bash
# Open scripts/api-test.http file
# Click "Send Request" to execute tests
```

### 2. Bash Script Testing

For Linux/Mac users:

```bash
chmod +x scripts/api-test.sh
./scripts/api-test.sh
```

For Windows users (Git Bash):

```bash
bash scripts/api-test.sh
```

## ⚙️ Configuration

### Configuration Hierarchy

1. `config/default.yaml` - Base configuration (shared across all environments)
2. `config/{env}.yaml` - Environment-specific configuration (overrides defaults)

### Configuration Options

```yaml
# Server configuration
server:
  port: "8080"              # Server port
  mode: "debug"             # Running mode: debug, release, test

# Database configuration
database:
  host: "localhost"         # Database host
  port: "3306"              # Database port
  user: "root"              # Database username
  password: ""              # Database password
  dbname: "gin_boilerplate" # Database name

# JWT configuration
jwt:
  secret: "your-secret-key" # JWT secret key (MUST change in production)
  expire_time: 24           # Token validity period (hours)
```

### Custom Startup Banner

Edit the `config/banner.txt` file to customize your startup banner.

## 🔧 Development Guide

### Adding New APIs

1. **Create Model** (`models/`)

```go
type Product struct {
    BaseModel
    Name  string `gorm:"not null" json:"name"`
    Price float64 `json:"price"`
}
```

2. **Create Service** (`services/`)

```go
type ProductService struct{}

func (s *ProductService) CreateProduct(product *models.Product) error {
    return database.GetDB().Create(product).Error
}
```

3. **Create Controller** (`controllers/`)

```go
type ProductController struct {
    productService *services.ProductService
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
    // Handle request
}
```

4. **Register Routes** (`router/router.go`)

```go
productController := controllers.NewProductController()
productRoutes := authenticated.Group("/products")
{
    productRoutes.POST("", productController.CreateProduct)
    productRoutes.GET("", productController.GetAllProducts)
}
```

### Using Middleware

```go
// Global middleware
r.Use(middleware.Logger())

// Route group middleware
authenticated := v1.Group("")
authenticated.Use(middleware.JWTAuth())
```

### Database Migration

Add auto-migration for new models in `main.go`:

```go
database.GetDB().AutoMigrate(
    &models.User{},
    &models.Product{}, // New model
)
```

## 🛡️ Security Recommendations

1. **Change JWT Secret**: Use a strong secret key in production
2. **HTTPS**: Use HTTPS in production environments
3. **Database Password**: Don't commit production config files to Git
4. **Input Validation**: Use Gin's binding for user input validation
5. **Rate Limiting**: Add API rate limiting middleware as needed

## 📦 Dependencies

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Viper](https://github.com/spf13/viper) - Configuration management
- [JWT](https://github.com/golang-jwt/jwt) - JWT authentication
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Password encryption

## 📝 TODO

- [ ] Add unit tests
- [ ] Add API documentation (Swagger)
- [x] Add Docker support
- [ ] Add rate limiting middleware
- [ ] Add cache support (Redis)
- [ ] Add file logging

## 📄 License

MIT License

## 🤝 Contributing

Issues and Pull Requests are welcome!

---

**Happy Coding!** 🎉
