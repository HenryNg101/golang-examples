# Golang Backend Playground

A hands-on playground for exploring backend concepts using **Golang**.

This repo is a collection of small, focused demos and code snippets covering things like:

* Core Go programming patterns
* PostgreSQL with GORM
* Elasticsearch search APIs
* gRPC contracts and services
* Redis usage
* Docker-based local infrastructure

The goal is simple: **learn by building small, real pieces of backend systems**.

---

## 📦 Project Structure

```
.
├── docker-compose.yml
├── .env
├── elasticsearch/     # Elasticsearch + gRPC demos
├── sql_db/            # PostgreSQL + GORM demos
├── programming/       # General Go snippets & experiments
├── go.work            # Go workspace file
```

Each top-level folder is its **own Go module**, with multiple packages inside.

Executable programs live inside `/cmd` directories.

---

## 🚀 Getting Started

### 1. Clone & Setup Environment

Create a `.env` file in the root. To setup different parts of the project, use these sets of env variables:

- Postgres:
    ```
    POSTGRES_USER=...
    POSTGRES_PASSWORD=...
    POSTGRES_DB=...
    POSTGRES_HOST=localhost
    POSTGRES_PORT=5432
    ```

- Elasticsearch + Kibana:
    ```
    ELASTIC_USER=elastic
    ELASTIC_PASSWORD=...
    KIBANA_PASSWORD=...
    ELASTIC_HOST=localhost
    ELASTIC_PORT=9200
    ELASTIC_DATA_STREAM_SOURCE=...
    ```

- Redis:
    ```
    REDIS_PASSWORD=...
    ```

You can use one, or some, or all of them altogether. It depends on what service(s) you need to run

---

### 2. Start Infrastructure

This project uses Docker Compose to spin up dependencies:

* PostgreSQL
* Elasticsearch
* Kibana
* Redis

Run:

```
docker compose up
```

---

### 3. Access Services

* **PostgreSQL** → `localhost:5432`
* **Elasticsearch** → `http://localhost:9200`
* **Kibana** → `http://localhost:5601`
* **Redis** → `localhost:6379` for Redis server, `localhost:8001` for RedisInsight GUI tool for Redis

---

## ▶️ Running Code

Each module contains runnable programs inside `/cmd`.

### Examples:

Run PostgreSQL CRUD demo:

```
go run ./sql_db/cmd/crud
```

Run Elasticsearch search demo:

```
go run ./elasticsearch/cmd/search
```

> Each `/cmd` folder contains a `main.go` entry point.

---

## 🧩 Modules Overview

### 📊 `/sql_db`

PostgreSQL + GORM experiments

Includes:

* Basic CRUD operations
* Relational modeling
* Query patterns

👉 See [`/sql_db/README.md`](./sql_db/README.md) for schema & details

---

### 🔍 `/elasticsearch`

Elasticsearch + Kibana + gRPC demos

Includes:

* Indexing & searching
* Query DSL examples
* gRPC service integration

👉 See [`/elasticsearch/README.md`](./elasticsearch/README.md)

---

### 🧪 `/programming`

General Go experiments

Includes:

* Language features
* Patterns & utilities
* Small isolated snippets

---

## 🗄️ Infrastructure Overview

The `docker-compose.yml` sets up:

* **Elasticsearch (single-node)** with authentication
* **Kibana** connected via service account
* **PostgreSQL**
* **Redis (with basic password protection)**

This setup is intentionally **minimal and dev-only**:

* No TLS
* No clustering
* No production hardening

---

## 🔄 Reset Everything

To wipe all data and start fresh:

```
docker compose down -v
```

---

## ⚠️ Notes

* This is a **learning project**, not production-ready code
* Security is minimal (no HTTPS, basic auth only)
* Expect rough edges and experimentation

---

## 🧭 Philosophy

Instead of building one big system, this repo focuses on:

> **Small, isolated, testable backend concepts**

Each folder answers questions like:

* “How do I model this in SQL?”
* “How does Elasticsearch actually query this?”
* “What does this look like in Go?”

---

## ✍️ Future Additions

Planned expansions:

* Message queues (Kafka)
* Caching strategies (Redis)
* Design patterns
* ...