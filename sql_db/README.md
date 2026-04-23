# 🗄️ PostgreSQL + GORM Demo

This module demonstrates how to work with **PostgreSQL** in Go using **GORM**.

It focuses on:
- Relational schema design
- CRUD operations
- Associations (1-to-many relationships)
- Query patterns in Go

---

## 🚀 How to Run

Make sure the root `docker-compose` is running:

```bash
docker compose up
````

Then run any executable inside `/cmd`:

```bash
go run ./sql_db/cmd/crud
```

---

## 🧱 Database Schema

### Table `users`

| Column     | Type        | Notes              |
| ---------- | ----------- | ------------------ |
| id         | bigserial   | Primary key        |
| name       | text        | User name          |
| created_at | timestamptz | Creation timestamp |

---

### Table `orders`

| Column     | Type        | Notes              |
| ---------- | ----------- | ------------------ |
| id         | bigserial   | Primary key        |
| user_id    | int8        | FK → users.id      |
| price      | int8        | Total order price  |
| created_at | timestamptz | Creation timestamp |

---

### Table `order_items`

| Column     | Type        | Notes              |
| ---------- | ----------- | ------------------ |
| id         | bigserial   | Primary key        |
| order_id   | int8        | FK → orders.id     |
| price      | int8        | Item price         |
| name       | text        | Item name          |
| created_at | timestamptz | Creation timestamp |

---

## 🔗 Relationships

* A **user** has many **orders**
* An **order** has many **order_items**

```
users → orders → order_items
```

---

## 🧪 What This Module Covers

* Mapping Go structs to SQL tables using GORM
* Handling foreign key relationships
* Performing:

  * Inserts
  * Queries (with joins / preloading)
  * Updates
  * Deletes
* Structuring database code in a clean way

---

## ⚠️ Notes

* This is a **learning/demo setup**
* Schema is intentionally simple
* No migrations tooling included (yet)

---

## 📌 Ideas for Expansion

* Add migrations (golang-migrate / atlas)
* Add transactions examples
* Benchmark query approaches