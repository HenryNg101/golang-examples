# Elasticsearch + Kibana (Docker Compose)

Minimal local setup for running a single-node Elasticsearch cluster with Kibana, using basic authentication (no TLS).

## 🚀 Quick Start

1. Create a `.env` file in the same directory, set the password as you wish:

```
ELASTIC_PASSWORD=...
KIBANA_PASSWORD=...
```

2. Start the stack:

```
docker compose up
```

3. Open Kibana:

```
http://localhost:5601
```

4. Log in with:

* **Username:** `elastic`
* **Password:** `${ELASTIC_PASSWORD}` (from `.env`)

---

## ⚙️ What this setup does

* Runs Elasticsearch with security enabled
* Automatically sets the `kibana_system` user password via API
* Connects Kibana using that service account
* Avoids TLS/cert complexity (dev-only)

---

## 🧹 Reset / Fresh Start

To reset everything (including data volume):

```
docker compose down -v
```

---

## ⚠️ Notes

* This setup is **for local development only**
* No HTTPS / TLS is configured
* Do not use in production as-is

---
