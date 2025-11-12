# Microservice B - Sensor Data Storage & API

## Overview

Microservice B receives sensor data from Microservice A (via gRPC/MQTT), stores it in a MySQL database, and exposes a REST API for data retrieval, update, and deletion. It supports authentication via JWT.

---

## Features

* Receive and store sensor data:

  ```json
  {
    "sensor_value": 12.34,
    "sensor_type": "temperature",
    "ID1": "A",
    "ID2": 1,
    "timestamp": "2025-11-12T12:00:00Z"
  }
  ```
* Retrieve data by:

  * Combination of IDs (`ID1` & `ID2`)
  * Timestamp/duration range
  * Both IDs and timestamps
* Update and delete data by filters
* Pagination for data retrieval
* JWT authentication for all `/data` endpoints
* Scalable to handle multiple Microservice A instances

---

## Tech Stack

* Language: Go
* Framework: Echo
* Database: MySQL
* Authentication: JWT
* Architecture: Microservices

---

## Installation

1. Clone the repository:

```bash
git clone <repo-url>
cd microservice-b
```

2. Configure MySQL connection in `.env`:

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=sensordb
```

3. Run the service:

```bash
go run ./cmd
```

---

## Database Schema

**Table: sensor_data**

| Column       | Type        | Notes                       |
| ------------ | ----------- | --------------------------- |
| id           | bigint      | primary key, auto_increment |
| sensor_value | double      |                             |
| sensor_type  | varchar(50) |                             |
| id1          | char(1)     |                             |
| id2          | int         |                             |
| timestamp    | datetime    |                             |

---

## API Endpoints

### 1. Login

```
POST /login
```

**Request:**

```json
{
  "username": "admin",
  "password": "password123"
}
```

**Response:**

```json
{
  "token": "<JWT_TOKEN>"
}
```

---

### 2. Retrieve Data

```
GET /data
```

**Query Params:**

* `ID1` - optional
* `ID2` - optional
* `start` - optional, RFC3339 timestamp
* `end` - optional, RFC3339 timestamp
* `page` - optional, default 1
* `limit` - optional, default 20

**Example:**

```bash
curl -H "Authorization: Bearer <TOKEN>" \
"http://localhost:8081/data?ID1=A&ID2=1&start=2025-11-12T00:00:00Z&end=2025-11-12T12:00:00Z&page=1&limit=10"
```

---

### 3. Update Data

```
PUT /data
```

**Query Params:** `ID1`, `ID2`, `start`, `end`, `value`

**Example:**

```bash
curl -X PUT -H "Authorization: Bearer <TOKEN>" \
"http://localhost:8081/data?ID1=A&ID2=1&start=2025-11-12T00:00:00Z&end=2025-11-12T12:00:00Z&value=25.5"
```

---

### 4. Delete Data

```
DELETE /data
```

**Query Params:** `ID1`, `ID2`, `start`, `end`

**Example:**

```bash
curl -X DELETE -H "Authorization: Bearer <TOKEN>" \
"http://localhost:8081/data?ID1=A&ID2=1&start=2025-11-12T00:00:00Z&end=2025-11-12T12:00:00Z"
```

---

## Usage Notes

* Use `/login` to get a JWT token before accessing `/data` endpoints.
* Supports filtering by IDs, timestamps, or both.
* Pagination allows handling large datasets efficiently.
* Scalable to handle multiple Microservice A instances sending sensor data.
