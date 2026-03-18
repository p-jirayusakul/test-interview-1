# test-interview-1

API สำหรับจัดการสินค้า พัฒนาด้วย Go, Gin, PostgreSQL และ Docker

## Tech Stack
- Go 1.26
- Gin
- PostgreSQL
- Docker / Docker Compose


## How to run with Docker

### 1. Start services
```bash
bash docker compose up --build
```

### 2. API endpoint
โดยค่าเริ่มต้น API จะรันที่:
```text
http://localhost:8080/api/v1
```

## API Documentation

Swagger UI:
```text
http://localhost:8080/swagger/index.html
```

## Testing

### Run unit tests
```bash
go test ./...
```

### Run E2E tests
```bash
go test ./tests/e2e -v
```

## Notes
- โปรเจกต์นี้แยก layer ตามแนวคิดคล้าย Clean Architecture
- ใช้ PostgreSQL เป็นฐานข้อมูลหลัก
- ใช้ Docker Compose สำหรับรันระบบทั้งหมดแบบพร้อมใช้งาน