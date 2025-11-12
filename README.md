Perfect! Let’s create a **combined README** that explains **Microservice A and B together**, including their interaction, architecture, and usage. I’ll keep it clear, simple, and visual-friendly in Markdown.

---

# Sensor Microservices System

## Overview

This system consists of **two microservices** that simulate, store, and manage sensor data:

1. **Microservice A** – Sensor Data Generator
2. **Microservice B** – Sensor Data Storage & API

The services interact in a **microservice architecture**:

* Microservice A generates sensor data.
* Microservice B receives data (via gRPC/MQTT) and stores it in MySQL.
* Users can query, update, or delete stored data via REST APIs on Microservice B.

---

## Architecture

```
+-------------------+          gRPC/MQTT          +-------------------+
|                   | --------------------------> |                   |
| Microservice A    |                             | Microservice B    |
| (Sensor Generator)|                             | (Storage & API)   |
|                   |                             |                   |
+-------------------+                             +-------------------+
          ^                                               ^
          |                                               |
          | REST API to control frequency                 | REST API for users
          v                                               |
    (Update frequency)                                    |
                                                          (JWT Auth)
```

* **Microservice A**: Generates random sensor data with fields: `sensor_value`, `sensor_type`, `ID1`, `ID2`, `timestamp`.
* **Microservice B**: Receives data, stores in MySQL, and provides REST API with JWT authentication for data management.

---

## Tech Stack

| Component        | Technology    |
| ---------------- | ------------- |
| Language         | Go            |
| Web Framework    | Echo          |
| Database         | MySQL         |
| Auth             | JWT           |
| Communication    | gRPC   |
| Architecture     | Microservices |
| Containerization | Docker        |

---

## Microservice A

**Purpose**: Generate random sensor data and optionally control the frequency.

### API

* `GET /status` – check service status
* `PUT /frequency` – update data generation frequency (ms)

**Run Multiple Instances**:

```bash
./microservice-a --sensor-type=temperature
./microservice-a --sensor-type=humidity
```

---

## Microservice B

**Purpose**: Receive sensor data, store in MySQL, and provide REST API for querying, updating, and deleting data.

### Database Schema

**Table: sensor_data**

| Column       | Type        |
| ------------ | ----------- |
| id           | bigint      |
| sensor_value | double      |
| sensor_type  | varchar(50) |
| id1          | char(1)     |
| id2          | int         |
| timestamp    | datetime    |

### API Endpoints (JWT Protected)

* `POST /login` – get JWT token (`admin/password123`)
* `GET /data` – retrieve data (filter by IDs, timestamps, or both)
* `PUT /data` – update sensor values
* `DELETE /data` – delete sensor values

**Example:**

```bash
curl -H "Authorization: Bearer <TOKEN>" \
"http://localhost:8081/data?ID1=A&ID2=1&start=2025-11-12T00:00:00Z&end=2025-11-12T12:00:00Z"
```

---

## Interaction Workflow

1. **Start Microservice B** (Database ready).
2. **Start one or more instances of Microservice A** for different sensor types.
3. Microservice A generates sensor data and sends it to Microservice B.
4. Users query, update, or delete sensor data from Microservice B via REST API with JWT authentication.
5. Frequency of data generation in Microservice A can be updated dynamically via REST.

---

## Usage Notes

* JWT token is required for all `/data` endpoints.
* Microservice B supports pagination and filtering for large datasets.
* The system can scale horizontally by running multiple instances of Microservice A.

---

## Quick Start

1. **Start MySQL** and create `sensordb` database.
2. **Run Microservice B**:

```bash
go run ./microservice-b/cmd
```

3. **Run Microservice A**:

```bash
go run ./microservice-a/cmd --sensor-type=temperature
```

4. **Login** to get JWT:

```bash
curl -X POST http://localhost:8081/login \
-H "Content-Type: application/json" \
-d '{"username":"admin","password":"password123"}'
```

**Access data endpoints** using the token.