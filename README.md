# E-Commerce API

A robust and scalable e-commerce API built with Go, featuring user authentication, product management, shopping cart functionality, and secure payment processing with Stripe integration.

## 🚀 Features

- **User Management**
    - JWT-based authentication
    - Role-based access control (Admin/User)
    - Secure password handling

- **Product Management**
    - Product creation and management (Admin)
    - Product search and filtering
    - Inventory tracking

- **Shopping Cart**
    - Add/remove products
    - Update quantities
    - Cart persistence

- **Order Processing**
    - Order creation and management
    - Multiple order statuses
    - Order history

- **Payment Integration**
    - Secure payment processing with Stripe
    - Payment status tracking
    - Refund handling

- **Admin Dashboard**
    - Product management
    - Order management
    - User management
    - Inventory control

## 🛠️ Technical Stack

- **Backend**: Go 1.21+
- **Framework**: Gin-Gonic
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT
- **Payment**: Stripe
- **Documentation**: Swagger/OpenAPI
- **Containerization**: Docker

## 📋 Prerequisites

- Go 1.21 or higher
- PostgreSQL 13 or higher
- Docker & Docker Compose (optional)
- Stripe Account (for payment processing)

## 🔧 Installation

1. Clone the repository:
```bash
git clone https://github.com/letsmakecakes/ecommerce-api
cd ecommerce-api
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run database migrations:
```bash
go run cmd/migrate/main.go up
```

5. Start the server:
```bash
go run cmd/api/main.go
```

## 🐳 Docker Setup

1. Build and run with Docker Compose:
```bash
docker-compose up --build
```

2. The API will be available at `http://localhost:8080`

## 🔑 Environment Variables

```env
# Server Configuration
SERVER_PORT=8080
SERVER_ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=ecommerce
DB_USER=postgres
DB_PASSWORD=your_password

# JWT Configuration
JWT_SECRET=your_jwt_secret
JWT_EXPIRY=24h

# Stripe Configuration
STRIPE_SECRET_KEY=your_stripe_secret_key
STRIPE_WEBHOOK_SECRET=your_stripe_webhook_secret
```

## 📚 API Documentation

API documentation is available at `/swagger/index.html` when running in development mode.

### Key Endpoints

```
POST   /api/v1/auth/register     # Register new user
POST   /api/v1/auth/login        # Login user
GET    /api/v1/products          # List products
POST   /api/v1/cart/items        # Add item to cart
POST   /api/v1/orders            # Create order
POST   /api/v1/orders/:id/pay    # Process payment

# Admin Routes
POST   /api/v1/admin/products    # Create product
PUT    /api/v1/admin/products/:id # Update product
GET    /api/v1/admin/orders      # List all orders
```

## 🧪 Testing

Run all tests:
```bash
go test ./...
```

Run specific test:
```bash
go test ./internal/domain/services -v
```

## 📦 Project Structure

```
📦 ecommerce-api
├── cmd/                  # Application entrypoints
├── internal/             # Private application code
│   ├── api/             # API layer
│   ├── domain/          # Business logic
│   ├── platform/        # External services
│   └── repository/      # Data access
├── pkg/                 # Public libraries
├── migrations/          # Database migrations
└── docs/               # Documentation
```

## 🔐 Security

- All passwords are hashed using bcrypt
- JWT tokens are used for authentication
- HTTPS enforced in production
- Input validation on all endpoints
- Role-based access control
- Rate limiting on sensitive endpoints

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- Your Name - [@letsmakecakes](https://github.com/letsmakecakes)

## 🙏 Acknowledgments

- [Gin-Gonic](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- [Stripe Go Library](https://github.com/stripe/stripe-go)