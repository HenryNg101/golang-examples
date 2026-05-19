# Kafka + Go Experiments

This repository is a hands-on playground to understand core Kafka concepts using simple Go programs (`segmentio/kafka-go`) and a local Docker-based Kafka setup.

---

# 🚀 Getting Started

## 1. Start Kafka

Run this on parent directory of the repo:

```bash
docker-compose up -d
```

Wait ~10–15 seconds for Kafka to fully initialize.

---

## 2. Create Topics

```bash
./scripts/topic-creation.sh
```

Creates:

* `orders` (3 partitions)
* `payments` (3 partitions)

---

## 3. Run Producer

```bash
go run ./cmd/producer
```

This will continuously produce messages to:

* `orders`
* `payments`

---

## 4. Run Consumers

Open separate terminals for each experiment below.


# 🧪 Experiments

## 1. Produce to Multiple Topics

**What to run:**

```bash
go run ./cmd/producer
```

**What it does:**

* Sends `order-*` messages → `orders`
* Sends `payment-*` messages → `payments`
* Uses keys (`i % 3`) to influence partitioning

**What to observe (Redpanda UI):**

* Messages distributed across partitions
* Same key → same partition

---

## 2. Consumer WITHOUT Group

**What to run:**

```bash
go run ./cmd/consumers/no-group
```

Run multiple instances.

**What happens:**

* Each consumer reads independently
* No coordination
* Duplicate processing across instances

**Key insight:**

```
No group = no load balancing
```

---

## 3. Consumer WITH Group

**What to run:**

```bash
go run ./cmd/consumers/consumer-group
```

Run multiple instances.

**What happens:**

* Kafka assigns partitions across consumers
* Each partition → only ONE consumer in group

**Try this:**

* Run 2 consumers → partitions split
* Run 4 consumers → one stays idle

**Key insight:**

```
Max parallelism = number of partitions
```

---

## 4. Multi-topic Consumer

**What to run:**

```bash
go run ./cmd/consumers/multi-topics
```

**What happens:**

* One consumer group reads from:

  * `orders`
  * `payments`

**Observe:**

* Messages interleaved from multiple topics

**Key insight:**

```
One consumer group can consume multiple topics
```

---

## 5. Partition Assignment & Rebalancing

**Steps:**

1. Start 2 consumers (consumer-group)
2. Start producer
3. Start a 3rd consumer

**Observe:**

* Partitions get reassigned
* Temporary pause during rebalance

**Then:**

* Kill one consumer
* Watch reassignment

**Key insight:**

```
Consumer group = dynamic load balancer
```

---

## 6. No Commit (Manual Offset Control)

**What to run:**

```bash
go run ./cmd/consumers/no-commit
```

**What happens:**

* Messages are read using `FetchMessage()`
* Offsets are NOT committed

**Try this:**

1. Run consumer
2. Stop it
3. Run again

**Observe:**

* Same messages are reprocessed

**Key insight:**

```
No commit → Kafka thinks message was never processed
```

---

## 7. Offset Reset

**What to run:**

```bash
./scripts/offset-reset.sh <group> <topic>
```

Example:

```bash
./scripts/offset-reset.sh order-processors orders
```

**What it does:**

* Shows what reset would do (`--dry-run`)
* Resets offsets to earliest

**Observe:**

* Consumer reprocesses all messages

**Key insight:**

```
Offsets control where consumption starts
```

---

# 🔍 Observing with Redpanda Console / Kafka UI

Use a UI tool (recommended):

## What to check:

### 1. Topics

* `orders`, `payments`
* Partition count

### 2. Messages

* Keys
* Partition distribution

### 3. Consumer Groups

* `order-processors`
* `message-processors`

### 4. Consumer Lag

* See how far behind consumers are

### 5. Partition Assignment

* Which consumer owns which partition

---

# 🧠 Key Concepts Demonstrated

## 1. Partitioning

* Messages with same key go to same partition

## 2. Ordering

* Guaranteed only within a partition

## 3. Consumer Groups

* Enable parallelism and load balancing

## 4. Offsets

* Track progress per group

## 5. At-least-once Delivery

* Messages may be reprocessed

## 6. Rebalancing

* Happens when consumers join/leave

---

# ⚠️ Notes

* This setup uses:

  * Single broker
  * Replication factor = 1

👉 Not production-safe (no fault tolerance)


# 💡 Planned Next Experiments (for future extensions)

* Add manual commit after processing
* Introduce artificial failures in consumer
* Add retry logic in producer
* Simulate slow consumers (lag)
* Build a pipeline: `orders → payments → notifications`

---

# 🧾 Summary

This repo demonstrates:

* Producing to multiple topics
* Consumer groups vs no groups
* Multi-topic consumption
* Partition-based parallelism
* Offset management (auto vs manual)
* Reprocessing via offset reset

Kafka becomes much easier once you *see* these behaviors live—so run multiple terminals, break things, and observe how the system reacts.
