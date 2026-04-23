# 🔍 Elasticsearch + Kibana + gRPC Demo

This module demonstrates how to work with **Elasticsearch** using Go, including:

- Search queries (Query DSL)
- Data ingestion (bulk indexing)
- Working with **data streams vs indices**
- Aggregation queries
- gRPC integration with search services

---

## 🚀 How to Run

From the project root:

```bash
docker compose up
````

Wait until:

* Elasticsearch is healthy
* Kibana is accessible

---

## 🔑 Access Kibana

Open:

```
http://localhost:5601
```

Login with:

* **Username:** `elastic`
* **Password:** `${ELASTIC_PASSWORD}` (from `.env`)

---

## 📥 Load Sample Dataset (Required)

This project uses **Kibana's sample web logs dataset**.

### Steps:

1. Open Kibana (`localhost:5601`)
2. Open the side menu
3. Click **"Add Integrations"**
4. Search for: `Sample Data`
5. Click into it
6. Go to:
   * **"Add data" → "Sample data"**
   * Then **"Other sample data sets"**
7. Install **"Sample web logs"**

---

## ⚙️ Configure Environment

After installing the dataset, set:

```env
ELASTIC_DATA_STREAM_SOURCE=kibana_sample_data_logs
```

---

## 📊 About the Dataset

This dataset represents **web server logs**, where each document is a request information

| Field        | Meaning                |
| ------------ | ---------------------- |
| `@timestamp` | When request happened  |
| `clientip`   | IP of the requester    |
| `request`    | Endpoint accessed      |
| `response`   | HTTP status code       |
| `bytes`      | Response size          |
| `geo.src`    | Country of origin      |
| `url`        | Full URL               |
| `agent`      | Browser/client info    |
| `referer`    | Traffic source         |
| `tags`       | success / error labels |
| `machine.os` | Client OS              |
| `extension`  | File type requested    |

---

## 🧠 What This Module Demonstrates

### 1. Working with Data Streams

* Source data comes from:

  ```
  kibana_sample_data_logs
  ```
* Data streams are:

  * Append-only
  * Not suitable for direct CRUD

---

### 2. Reindexing into Custom Index

This project demonstrates:

* Reading from a data stream
* Transforming documents
* Writing into a **custom index**

Why?

* Enables **full CRUD operations**
* More flexible mappings
* Better for application-level control

---

### 3. Bulk Ingestion

* Efficient indexing using `_bulk`
* Batching strategies

---

### 4. Search APIs

* Query DSL usage
* Filtering, matching, aggregations

---

### 5. gRPC Integration

* Exposing search functionality via gRPC
* Structuring services cleanly

---

## ▶️ Run Example

```bash
go run ./elasticsearch/cmd/search
```

---

## 🧹 Reset

To reset Elasticsearch data:

```bash
docker compose down -v
```

---

## ⚠️ Notes

* Single-node Elasticsearch (dev only)
* No TLS / HTTPS
* Minimal security setup
* Dataset is **provided by Kibana**, not custom

---

## 📌 Ideas for Expansion

* Add custom mappings & analyzers
* Compare keyword vs text fields
* Add pagination strategies (search_after)
* Introduce ranking/scoring tweaks