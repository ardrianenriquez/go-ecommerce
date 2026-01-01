

## Database

https://dbdiagram.io/d/Ecommerce-Go-693a1ac0e877c630745d8dc3

users:
  id, email, password_hash, role (customer, admin), timestamps

shops:
  id, name, description, status (open, close, suspended), timestamps

products:
  id, shop_id, name, description, sku, price_in_cents,
  max_order_quantity, stock_quantity, status, timestamps

carts:
  id, user_id, status (active, checked_out), expires_at, timestamps

cart_items:
  id, cart_id, product_id, quantity,
  unit_price_in_cents, total_price_in_cents, timestamps

orders:
  id, user_id, total_in_cents, status (pending_payment, paid, cancelled),
  payment_reference, timestamps

ordered_items:
  id, order_id, product_id (nullable),
  name, sku,
  quantity, unit_price_in_cents, total_price_in_cents,
  timestamps


users.email → unique index
carts.user_id, status → for active cart lookup
products.shop_id → for product listing per shop
orders.user_id

## Migration
https://betterstack.com/community/guides/scaling-go/golang-migrate/
https://github.com/golang-migrate/migrate/blob/v4.19.1/database/postgres/TUTORIAL.md

golang-migrate
- CLI for doing migration like rails g migration

curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xz
sudo mv migrate /usr/local/bin/

# Create migration
-ext = file extension
-dir = directory

migrate create -ext sql -dir migrations create_users_table

## Libraries
# Tidy and check
go mod tidy
go mod download

# Core web framework
go get github.com/gin-gonic/gin

# Database
go get github.com/lib/pq
go get github.com/golang-migrate/migrate/v4
go get github.com/golang-migrate/migrate/v4/database/postgres
go get github.com/golang-migrate/migrate/v4/source/file

# Config / Env
go get github.com/joho/godotenv

# Redis
go get github.com/redis/go-redis/v9

# Authentication (JWT)
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt

# Testing (add these for later)
go get github.com/stretchr/testify/assert
go get github.com/stretchr/testify/mock
go get github.com/testcontainers/testcontainers-go
go get github.com/testcontainers/testcontainers-go/modules/postgres
go get github.com/testcontainers/testcontainers-go/wait



### Initial plan
Go project structure (Gin)
A simple, clean structure inspired by common Go + Gin examples:

cmd/server/main.go

Wire config, DB, Redis, and Gin router, then router.Run().​

internal/config

Load env (DB URL, Redis, JWT secret, etc.).

internal/database

Postgres connection setup and migrations.

internal/models

Go structs for User, Shop, Product, Cart, CartItem, Order, OrderItem.

internal/repository

Interfaces and Postgres implementations: UserRepository, ShopRepository, ProductRepository, CartRepository, OrderRepository.

internal/service (business logic)

AuthService (signup/login, password hashing, JWT).

ShopService / ProductService (list shops/products, search, stock checks).

CartService (add/remove items, compute totals).

OrderService (checkout flow, create order in a transaction, reduce stock, clear cart).

internal/handler (Gin handlers)

Map Gin routes to service methods, handle request/response/validation.

pkg/api or internal/router

SetupRouter() that registers routes, groups, and middleware for Gin.​

Key Gin routes
Customer‑facing:

POST /v1/auth/signup, POST /v1/auth/login

GET /v1/shops, GET /v1/shops/:id/products

GET /v1/products (search/filter)

GET /v1/cart, POST /v1/cart/items, PATCH /v1/cart/items/:id, DELETE /v1/cart/items/:id

POST /v1/checkout → returns created order (status pending_payment or paid depending on how you simulate payment)

Admin:

GET /v1/admin/orders (filters: date range, shop, status)

GET /v1/admin/orders/:id

GET /v1/admin/reports/sales?from=...&to=...&shop_id=... (returns JSON and maybe CSV content)

Checkout design
In OrderService.Checkout(userID):

Load active cart with items for the user.

For each item, check product stock; fail if insufficient.

Calculate totals.

In a DB transaction:

Create orders row.

Create order_items.

Decrease products.stock.

Mark cart as checked_out.

Optionally push a job into Redis to send confirmation email.

This pattern is close to what many “Go + Gin e‑commerce” sample projects do, but tailored to your marketplace use case.​