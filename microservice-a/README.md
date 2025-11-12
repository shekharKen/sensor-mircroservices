# Microservice A - Sensor Data Generator

## Overview

Microservice A is responsible for **simulating sensor data generation** for various sensor types. It generates continuous data streams and allows for **dynamic adjustment of the generation frequency** via a simple REST API.

---

## Features

* **Sensor Data Generation:** Generates random sensor data following this JSON structure:
    ```json
    {
      "sensor_value": 12.34,
      "sensor_type": "temperature",
      "ID1": "A",
      "ID2": 1,
      "timestamp": "2025-11-12T12:00:00Z"
    }
    ```
* **Dynamic Frequency Control:** Update the data generation frequency in milliseconds via a REST API endpoint.
* **Multi-Instance Support:** Supports running multiple instances, with each instance configured for a **fixed sensor type** (e.g., one for temperature, one for humidity).
* **Logging:** Logs generated sensor data locally (planned for future transmission to Microservice B).

---

## Tech Stack

* **Language:** Go
* **Framework:** Echo
* **Architecture:** Microservices
* **Data Transport (Planned):** gRPC / MQTT (for integration with Microservice B)

---

## Installation & Run

1.  **Clone the repository:**
    ```bash
    git clone <repo-url>
    cd microservice-a
    ```

2.  **Build the service:**
    ```bash
    go build -o microservice-a ./cmd
    ```

3.  **Run the service:**
    ```bash
    ./microservice-a
    ```

---

## Usage Examples

### Starting Instances

You can start multiple instances, specifying the sensor type for each:

* **Temperature:**
    ```bash
    ./microservice-a --sensor-type=temperature
    ```
* **Humidity:**
    ```bash
    ./microservice-a --sensor-type=humidity
    ```

### Changing Generation Frequency

Use a `PUT` request to update the frequency dynamically (e.g., setting it to 1000ms):

```bash
curl -X PUT http://localhost:8080/frequency \
-H "Content-Type: application/json" \
-d '{"frequency_ms": 1000}'
```

## 🔍 Monitoring Data

Check your service logs for the generated output format:
```bash
Generated: {SensorValue:23.45 SensorType:temperature ID1:A ID2:1 Timestamp:2025-11-12 12:00:00}
```

## Notes

* **Future Integration:** Data is currently logged but is planned to be sent to **Microservice B** using **gRPC/MQTT**.
* **Scaling:** The architecture supports dynamic scaling by simply running additional instances.

