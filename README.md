# test-interview-1

API สำหรับจัดการสินค้า พัฒนาด้วย Go, Gin, PostgreSQL และ Docker

## Tech Stack
- Go 1.26
- Gin
- PostgreSQL
- Docker / Docker Compose

## Project Structure
```
test-interview-1
 cmd
  api
   main.go
 docs
  docs.go
  swagger.json
  swagger.yaml
 import
  migrations
   000001_init_schema.up.sql
 internal
  bootstrap
   server.go
  delivery
  domain
   product.go
   product_test.go
  infrastructure
   repository
    postgres
    product
     implement.go
     mapper.go
     model.go
  usecase
   product.go
   product_test.go
 pkg
 tests
 .env
 .gitignore
 docker-compose.yml
 Dockerfile
 go.mod
 README.md
External Libraries
Scratches and Consoles
```

## Clean Architecture
โปรเจกต์นี้แยกโค้ดตามแนวคิด Clean Architecture:
- `domain` เก็บ entity และ validation
- `usecase` เก็บ business flow
- `delivery` รับ request และส่ง response
- `infrastructure` คุยกับฐานข้อมูล
- `bootstrap` ทำหน้าที่ประกอบ dependency ทั้งระบบ

## How to run with Docker

### 1. Start services
```bash
docker compose up --build
```

### 2. API endpoint
โดยค่าเริ่มต้น API จะรันที่:
```text
http://localhost:8080/api/v1
```

## API Documentation

Swagger UI:
```text
http://localhost:8080/api-docs/index.html
```

### Generate Swagger Document
```bash
swag init --parseDependency --parseInternal -g cmd/api/main.go -o docs
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
- โปรเจกต์นี้แยก layer ตามแนวคิด Clean Architecture
- ใช้ PostgreSQL เป็นฐานข้อมูลหลัก
- ใช้ Docker Compose สำหรับรันระบบทั้งหมด
- มี unit test และ E2E test ครอบคลุมตาม requirement