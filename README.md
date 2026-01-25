# Tugas 1 - REST API di Golang

Aplikasi REST API sederhana untuk mengelola kategori film menggunakan Go.

## Apa yang dilakukan?

Aplikasi ini menyediakan API untuk CRUD (Create, Read, Update, Delete) kategori film. Data disimpan di memory, jadi akan hilang saat server di-restart.

## Fitur

- **GET** `/api/category` - Ambil semua kategori
- **POST** `/api/category` - Tambah kategori baru
- **GET** `/api/category/{id}` - Ambil kategori berdasarkan ID
- **PUT** `/api/category/{id}` - Update kategori berdasarkan ID
- **DELETE** `/api/category/{id}` - Hapus kategori berdasarkan ID

## Cara Menjalankan

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## Contoh Request

### Ambil semua kategori

```bash
curl http://localhost:8080/api/category
```

### Tambah kategori baru

```bash
curl -X POST http://localhost:8080/api/category \
  -H "Content-Type: application/json" \
  -d '{"name": "Comedy", "description": "Film lucu dan menghibur"}'
```

### Ambil kategori by ID

```bash
curl http://localhost:8080/api/category/1
```

### Update kategori

```bash
curl -X PUT http://localhost:8080/api/category/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Action Updated", "description": "Deskripsi baru"}'
```

### Hapus kategori

```bash
curl -X DELETE http://localhost:8080/api/category/1
```
