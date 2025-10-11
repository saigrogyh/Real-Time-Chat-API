#  Realtime Chat API (In-Progress)

This project is a **Realtime Chat Backend** built with **Go**, designed to demonstrate clean backend design and scalable structure for chat applications.

It follows **Clean Architecture** and **Hexagonal Architecture** principles — making the code easy to maintain, test, and extend.  
The backend currently supports **JWT authentication**, **user management**, and **chat room setup**, and it is ready for future WebSocket integration.

---

## ⚙️ Tech Stack

- **Go (Golang)** – main backend language  
- **Fiber** – fast, lightweight web framework  
- **GORM** – ORM for PostgreSQL  
- **PostgreSQL** – main database  
- **JWT** – authentication system  
- **Docker & docker-compose** – container setup for database and pgAdmin  

---

##  Project Structure

```
internal/
├── app/
│   ├── auth/          # Authentication and JWT handling
│   ├── user/          # User business logic
│   ├── chat/          # Chat room management
│   └── message/       # Message logic (for future websocket)
├── domain/            # Core entities (User, Chat, Message)
└── infrastructure/
    ├── db/            # Database connection (GORM + PostgreSQL)
    └── router/        # Fiber router and middleware setup

cmd/
└── main.go            # Application entry point
```

---

##  Features

- **User Registration & Login** with password hashing (`bcrypt`)  
- **JWT-based Authentication** for secure API access  
- **Chat Room Creation and Management** (user can create or join rooms)  
- **Scalable and Maintainable Structure** following clean architecture  
- Ready for **future WebSocket implementation** for real-time messaging  

---

##  Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/yorch/real-time-chat-api.git
cd real-time-chat-api
```

### 2. Create `.env` file

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=yorch
DB_PASSWORD=yorch123
DB_NAME=chatdb
JWT_SECRET=supersecretkey
APP_PORT=8080
```

### 3. Start PostgreSQL with Docker

```bash
docker-compose up -d
```

### 4. Run the application

```bash
go run cmd/main.go
```

The app will start at  
👉 **http://localhost:8080**

---

## 🧪 Example API Usage

### Register
```bash
POST /api/register
{
  "username": "yorch",
  "email": "yorch@example.com",
  "password": "1234"
}
```

### Login
```bash
POST /api/login
{
  "email": "yorch@example.com",
  "password": "1234"
}
```

### Create Chat Room
```bash
POST /api/chats
Authorization: Bearer <token>
{
  "name": "Golang Group"
}
```

### Get All Chat Rooms
```bash
GET /api/chats
Authorization: Bearer <token>
```

---

## 🧠 Future Improvements

- Add **WebSocket** for live messaging between users  
- Implement **message persistence** with timestamps  
- Add **user online/offline status** tracking  
- Add **unit tests** and integration tests  

---

## 🧑‍💻 Author

**Yorch**  
University Student · Backend Developer (Intern Candidate)  
Built with ❤️ using Go, Fiber, and PostgreSQL.


