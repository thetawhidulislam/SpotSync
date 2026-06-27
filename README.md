# 🚗 SpotSync - Smart Parking & EV Charging Reservation System

A RESTful backend API for managing smart parking zones and EV charging reservations. The system allows drivers to reserve parking spaces while ensuring that parking zones never exceed their available capacity through database transactions and row-level locking.

## 🌐 Live API

**Live URL:** https://spotsync-tawhid.up.railway.app

---

## ✨ Features

### Driver

* User Registration & Login (JWT Authentication)
* View all parking zones
* View parking zone details
* Reserve a parking/EV charging spot
* View personal reservations
* Cancel own reservation

### Admin

* Create parking zones
* Update parking zones
* Delete parking zones
* Set parking price per hour
* View all reservations

### Reservation System

* Transaction-safe reservation creation
* Row-Level Locking (`FOR UPDATE`)
* Prevents overbooking during concurrent requests
* Dynamic parking availability calculation

---

## 🛠 Tech Stack

### Backend

* Go (Golang 1.22+)
* Echo v5
* GORM

### Database

* PostgreSQL (Neon)

### Authentication

* JWT

### Deployment

* Railway

---

# 🏗 Project Architecture

The project follows a layered architecture.

```text
                 Client (Postman / Frontend)
                           │
                           ▼
                    HTTP Handlers
                           │
                           ▼
                      Service Layer
                           │
                           ▼
                   Repository Layer
                           │
                           ▼
                     PostgreSQL Database
```

### Layer Responsibilities

### Handler

* Receives HTTP requests
* Validates request
* Calls service methods
* Returns JSON response

### Service

* Contains business logic
* Handles validation
* Performs reservation rules
* Coordinates repositories

### Repository

* Handles database queries
* Uses GORM
* Performs transactions
* Implements Row-Level Locking

### Database

* Stores users
* Parking zones
* Reservations

---

# 🚨 Concurrency Protection

SpotSync prevents overbooking using:

* Database Transactions
* Row-Level Locking (`FOR UPDATE`)

When two users attempt to reserve the last parking space simultaneously:

1. First transaction locks the parking zone.
2. Second request waits.
3. Capacity is rechecked.
4. If full, reservation is rejected.

This guarantees that the total reservations never exceed the parking zone capacity.

---

# 📁 Project Structure

```text
cmd/
    main.go

internal/
    auth/
    config/
    domain/
        user/
        zone/
        reservation/
    middlewares/
    server/
    validator/

go.mod
README.md
```

---

# ⚙️ Setup

## Clone Repository

```bash
git clone <repository-url>

cd spotsync
```

---

## Install Dependencies

```bash
go mod download
```

---

## Environment Variables

Create a `.env` file.

```env
PORT=8080

DSN=postgresql://username:password@host/database?sslmode=require

JWT_SECRET=your_secret_key
```

---

## Run the Server

```bash
go run ./cmd
```

Server starts on:

```
http://localhost:8080
```

---

# 🚀 API Endpoints

## Authentication

| Method | Endpoint              |
| ------ | --------------------- |
| POST   | /api/v1/auth/register |
| POST   | /api/v1/auth/login    |

---

## Parking Zones

| Method | Endpoint          | Access |
| ------ | ----------------- | ------ |
| GET    | /api/v1/zones     | Public |
| GET    | /api/v1/zones/:id | Public |
| POST   | /api/v1/zones     | Admin  |
| PUT    | /api/v1/zones/:id | Admin  |
| DELETE | /api/v1/zones/:id | Admin  |

---

## Reservations

| Method | Endpoint                             | Access |
| ------ | ------------------------------------ | ------ |
| POST   | /api/v1/reservations                 | Driver |
| GET    | /api/v1/reservations/my-reservations | Driver |
| DELETE | /api/v1/reservations/:id             | Driver |
| GET    | /api/v1/reservations                 | Admin  |

---

# 🔐 Authentication

Protected endpoints require a JWT token.

Example:

```http
Authorization: Bearer <your_token>
```

---

# 📌 Reservation Flow

```text
User
   │
   ▼
Reserve Spot
   │
   ▼
Start Transaction
   │
   ▼
Lock Parking Zone (FOR UPDATE)
   │
   ▼
Check Capacity
   │
   ├───────────────► Full
   │                  │
   │                  ▼
   │             Reject Request
   │
   ▼
Create Reservation
   │
   ▼
Commit Transaction
```

---

# 📬 Example Success Response

```json
{
  "success": true,
  "message": "Reservation confirmed successfully",
  "data": {
    "id": 1,
    "license_plate": "ABC-1234",
    "status": "active"
  }
}
```

---

# 👨‍💻 Author

**Tawhidul Islam**

Backend Developer | Golang Developer

---
