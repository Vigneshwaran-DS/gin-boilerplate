# Gin Boilerplate

一个基于 Gin + GORM 的 Go Web 应用脚手架，采用最佳实践进行代码分层，开箱即用。

## ✨ 特性

- 🚀 **完整的项目结构** - 清晰的代码分层（Controller、Service、Model、Router）
- 🔐 **JWT 认证** - 完整的用户认证和授权体系
- ⚙️ **多环境配置** - 支持开发、生产等多环境配置（基于 Viper）
- 🗄️ **数据库 ORM** - 使用 GORM，支持自动迁移
- 🔒 **密码加密** - 使用 bcrypt 加密用户密码
- 📝 **日志中间件** - 请求日志记录
- 🌐 **CORS 支持** - 跨域资源共享中间件
- 📦 **统一响应格式** - 标准化 API 响应结构
- 🔄 **RESTful API** - 符合 REST 规范的 API 设计

## 📁 项目结构

```
gin-boilerplate/
├── config/                 # 配置文件
│   ├── config.go          # 配置加载逻辑
│   ├── default.yaml       # 默认配置
│   ├── development.yaml   # 开发环境配置
│   └── production.yaml    # 生产环境配置
├── controllers/           # 控制器层
│   ├── auth_controller.go
│   └── user_controller.go
├── database/              # 数据库连接
│   └── database.go
├── middleware/            # 中间件
│   ├── auth.go           # JWT 认证中间件
│   ├── cors.go           # CORS 中间件
│   └── logger.go         # 日志中间件
├── models/                # 数据模型层
│   ├── base.go           # 基础模型
│   └── user.go           # 用户模型
├── router/                # 路由层
│   └── router.go
├── scripts/               # 脚本文件
│   └── init.sql          # 数据库初始化脚本
├── services/              # 业务逻辑层
│   ├── auth_service.go
│   └── user_service.go
├── utils/                 # 工具类
│   ├── jwt.go            # JWT 工具
│   └── response.go       # 响应工具
├── .gitignore
├── go.mod
├── main.go               # 程序入口
└── README.md
```

## 🚀 快速开始

### 1. 环境要求

- Go 1.19+
- MySQL 5.7+

### 2. 克隆项目

```bash
git clone <repository-url>
cd gin-boilerplate
```

### 3. 安装依赖

```bash
go mod tidy
```

### 4. 配置数据库

#### 初始化数据库

```bash
mysql -u root -p < scripts/init.sql
```

#### 配置数据库连接

复制并修改开发环境配置：

```bash
cp config/development.yaml.example config/development.yaml
```

编辑 `config/development.yaml`：

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

### 5. 运行项目

#### 开发环境

```bash
go run main.go
# 或指定环境
go run main.go -e development
```

#### 生产环境

```bash
go run main.go -e production
```

服务默认运行在 `http://localhost:8080`

## 📚 API 文档

### 认证相关

#### 用户注册

```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "123456",
  "full_name": "Test User"
}
```

**响应：**

```json
{
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

#### 用户登录

```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "123456"
}
```

**响应：**

```json
{
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

### 用户相关（需要认证）

所有用户相关接口需要在 Header 中携带 Token：

```http
Authorization: Bearer {your_token}
```

#### 获取当前用户信息

```http
GET /api/v1/me
Authorization: Bearer {token}
```

#### 获取用户列表

```http
GET /api/v1/users?page=1&page_size=10
Authorization: Bearer {token}
```

#### 获取用户详情

```http
GET /api/v1/users/:id
Authorization: Bearer {token}
```

#### 更新用户

```http
PUT /api/v1/users/:id
Authorization: Bearer {token}
Content-Type: application/json

{
  "full_name": "Updated Name"
}
```

#### 删除用户

```http
DELETE /api/v1/users/:id
Authorization: Bearer {token}
```

### 健康检查

```http
GET /api/v1/health
```

**响应：**

```json
{
  "status": "ok",
  "message": "Service is running"
}
```

## ⚙️ 配置说明

### 配置文件层级

1. `config/default.yaml` - 基础配置（所有环境共享）
2. `config/{env}.yaml` - 环境特定配置（会覆盖默认配置）

### 配置项说明

```yaml
# 服务配置
server:
  port: "8080"              # 服务端口
  mode: "debug"             # 运行模式: debug, release, test

# 数据库配置
database:
  host: "localhost"         # 数据库地址
  port: "3306"              # 数据库端口
  user: "root"              # 数据库用户名
  password: ""              # 数据库密码
  dbname: "gin_boilerplate" # 数据库名

# JWT 配置
jwt:
  secret: "your-secret-key" # JWT 密钥（生产环境务必修改）
  expire_time: 24           # Token 有效期（小时）
```

## 🔧 开发指南

### 添加新的 API

1. **创建模型** (`models/`)

```go
type Product struct {
    BaseModel
    Name  string `gorm:"not null" json:"name"`
    Price float64 `json:"price"`
}
```

2. **创建服务** (`services/`)

```go
type ProductService struct{}

func (s *ProductService) CreateProduct(product *models.Product) error {
    return database.GetDB().Create(product).Error
}
```

3. **创建控制器** (`controllers/`)

```go
type ProductController struct {
    productService *services.ProductService
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
    // 处理请求
}
```

4. **注册路由** (`router/router.go`)

```go
productController := controllers.NewProductController()
productRoutes := authenticated.Group("/products")
{
    productRoutes.POST("", productController.CreateProduct)
    productRoutes.GET("", productController.GetAllProducts)
}
```

### 使用中间件

```go
// 全局中间件
r.Use(middleware.Logger())

// 路由组中间件
authenticated := v1.Group("")
authenticated.Use(middleware.JWTAuth())
```

### 数据库迁移

在 `main.go` 中添加新模型的自动迁移：

```go
database.GetDB().AutoMigrate(
    &models.User{},
    &models.Product{}, // 新增模型
)
```

## 🛡️ 安全建议

1. **修改 JWT Secret**：生产环境务必使用强密钥
2. **HTTPS**：生产环境使用 HTTPS
3. **数据库密码**：不要将生产环境配置文件提交到 Git
4. **输入验证**：使用 Gin 的 binding 验证用户输入
5. **限流**：根据需要添加 API 限流中间件

## 📦 依赖包

- [Gin](https://github.com/gin-gonic/gin) - Web 框架
- [GORM](https://gorm.io/) - ORM 库
- [Viper](https://github.com/spf13/viper) - 配置管理
- [JWT](https://github.com/golang-jwt/jwt) - JWT 认证
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - 密码加密

## 📝 TODO

- [ ] 添加单元测试
- [ ] 添加 API 文档（Swagger）
- [ ] 添加 Docker 支持
- [ ] 添加限流中间件
- [ ] 添加缓存支持（Redis）
- [ ] 添加日志文件输出

## 📄 License

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

---

**Happy Coding!** 🎉
