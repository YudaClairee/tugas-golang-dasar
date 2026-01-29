# Tugas 2 - REST API dengan Layered Architecture

Aplikasi REST API untuk mengelola kategori dan produk menggunakan Go dengan PostgreSQL database.

## Apa yang dilakukan?

Aplikasi ini udah direfactor jadi layered architecture yang proper dan menerapkan JOIN SQL untuk relasi antar tabel. Sekarang data disimpan persistent di PostgreSQL, bukan lagi di memory.

## Arsitektur

Proyek ini menggunakan **Layered Architecture** dengan struktur sebagai berikut:

```
.
├── models/           # Model data & struct
│   ├── category.go
│   └── product.go
├── repositories/     # Database access layer (SQL queries)
│   ├── category_repository.go
│   └── product_repository.go
├── services/         # Business logic layer
│   ├── category_service.go
│   └── product_service.go
├── handlers/         # HTTP handlers (routing & response)
│   ├── category_handler.go
│   └── product_handler.go
├── database/         # Database connection setup
│   └── database.go
└── main.go          # Entry point & dependency injection
```

### Flow:

`Handler → Service → Repository → Database`

### Kenapa Layered?

- **Separation of concerns** - setiap layer punya tanggung jawab yang jelas
- **Testability** - mudah di-mock dan di-test per layer
- **Maintainability** - gampang di-maintain dan dikembangin
- **Reusability** - logic bisa dipake ulang

## Fitur JOIN SQL

Implementasi **LEFT JOIN** antara tabel `product` dan `category`:

```sql
SELECT p.id, p.name, p.price, p.stock, p.category_id,
       c.id, c.name, c.description
FROM product p
LEFT JOIN category c ON p.category_id = c.id
```

JOIN ini bikin pas ambil data produk, langsung dapet info kategorinya juga dalam satu query.

## Database

Pakai PostgreSQL dengan 2 tabel utama:

### Table: category

- `id` (serial, primary key)
- `name` (varchar)
- `description` (text)

### Table: product

- `id` (serial, primary key)
- `name` (varchar)
- `price` (numeric)
- `stock` (integer)
- `category_id` (integer, foreign key ke category.id)

## Environment Variables

Buat file `.env` di root project:

```env
PORT=8080
DB_CONN=postgres://username:password@localhost:5432/dbname?sslmode=disable
```

## Cara Menjalankan

```bash
# Install dependencies
go mod download

# Run server
go run main.go
```

Server akan jalan di `http://localhost:8080`

## API Endpoints

### Category Endpoints

- **GET** `/api/category` - Ambil semua kategori
- **POST** `/api/category` - Tambah kategori baru
- **GET** `/api/category/{id}` - Ambil kategori by ID
- **PUT** `/api/category/{id}` - Update kategori
- **DELETE** `/api/category/{id}` - Hapus kategori

### Product Endpoints

- **GET** `/api/product` - Ambil semua produk (dengan info kategori via JOIN)
- **POST** `/api/product` - Tambah produk baru
- **GET** `/api/product/{id}` - Ambil produk by ID (dengan info kategori via JOIN)
- **PUT** `/api/product/{id}` - Update produk
- **DELETE** `/api/product/{id}` - Hapus produk

## Contoh Request

### Ambil semua produk (dengan kategori)

```bash
curl http://localhost:8080/api/product
```

Response:

```json
[
  {
    "id": 1,
    "name": "Spiderman Action Figure",
    "price": 150000,
    "stock": 10,
    "category_id": 1,
    "category": {
      "id": 1,
      "name": "Action",
      "description": "Film aksi spektakuler"
    }
  }
]
```

### Tambah produk baru

```bash
curl -X POST http://localhost:8080/api/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Batman Figure",
    "price": 200000,
    "stock": 5,
    "category_id": 1
  }'
```

### Tambah kategori baru

```bash
curl -X POST http://localhost:8080/api/category \
  -H "Content-Type: application/json" \
  -d '{"name": "Comedy", "description": "Film lucu dan menghibur"}'
```

### Update produk

```bash
curl -X PUT http://localhost:8080/api/product/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Spiderman Deluxe Figure",
    "price": 175000,
    "stock": 8,
    "category_id": 1
  }'
```

### Hapus produk

```bash
curl -X DELETE http://localhost:8080/api/product/1
```

## Dependencies

- `github.com/lib/pq` - PostgreSQL driver
- `github.com/spf13/viper` - Configuration management

## What's New di Tugas 2?

✅ Layered Architecture (Handler → Service → Repository)  
✅ PostgreSQL integration  
✅ LEFT JOIN untuk relasi product-category  
✅ Dependency injection pattern  
✅ Environment variable dengan viper  
✅ Connection pooling  
✅ CRUD lengkap untuk Product & Category
