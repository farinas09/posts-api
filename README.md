# Go REST API

A modern REST API built with Go that provides JWT authentication and posts management.

## 🚀 Features

- **JWT Authentication**: Complete registration and login system with JWT tokens
- **Post Management**: Creation and management of user posts
- **PostgreSQL Database**: Data persistence with PostgreSQL
- **Clean Architecture**: Clear separation between handlers, models, repositories and middleware
- **Docker Support**: Complete configuration with Docker Compose
- **Authentication Middleware**: Automatic JWT token validation
- **Password Hashing**: Security with bcrypt for passwords

## 📋 Requirements

- Go 1.24.3 or higher
- Docker and Docker Compose
- PostgreSQL (if not using Docker)

## 🛠️ Installation

### 1. Clone the repository

```bash
git clone <your-repository>
cd go-rest
```

### 2. Configure environment variables

Create a `.env` file in the project root:

```env
PORT=8080
JWT_SECRET=your_very_secure_jwt_secret_here
DATABASE_URL=postgres://postgres:postgres123@localhost:5432/go_rest_db?sslmode=disable
```

### 3. Start the database with Docker

```bash
cd database
docker-compose up -d
```

### 4. Install dependencies

```bash
go mod download
```

### 5. Run the application

```bash
go run main.go
```

The server will be available at `http://localhost:8080`

## 📚 API Endpoints

### Authentication

#### User Registration
```http
POST /signup
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "password123"
}
```

#### Login
```http
POST /login
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "password123"
}
```

**Response:**
```json
{
    "id": 1,
    "email": "user@example.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Current User Information
```http
GET /me
Authorization: <your_jwt_token>
```

### Posts

#### Create Post
```http
POST /posts
Authorization: <your_jwt_token>
Content-Type: application/json

{
    "title": "My first post",
    "content": "Post content here..."
}
```

#### Home Page
```http
GET /
```

## 🏗️ Project Structure

```
go-rest/
├── database/           # PostgreSQL configuration with Docker
│   ├── docker-compose.yml
│   ├── Dockerfile
│   ├── postgres.go
│   └── up.sql
├── handlers/           # HTTP controllers
│   ├── home.go
│   ├── post.go
│   └── user.go
├── middleware/         # Authentication middleware
│   └── auth.go
├── models/            # Data models
│   ├── claims.go
│   ├── post.go
│   └── user.go
├── repository/        # Data access layer
│   ├── post.go
│   └── user.go
├── server/           # Server configuration
│   └── server.go
├── main.go           # Entry point
├── go.mod
└── go.sum
```

## 🔧 Technologies Used

- **Go 1.24.3**: Programming language
- **Gorilla Mux**: HTTP router for Go
- **JWT**: Authentication with JSON Web Tokens
- **PostgreSQL**: Relational database
- **bcrypt**: Password hashing
- **Docker**: Containerization
- **godotenv**: Environment variable management

## 🔐 Security

- **JWT Tokens**: Stateless authentication with JWT tokens
- **Password Hashing**: Passwords are stored hashed with bcrypt
- **Validation Middleware**: Automatic token validation on protected routes
- **Environment Variables**: Sensitive configuration through environment variables

## 🐳 Docker

The project includes complete Docker configuration for PostgreSQL:

```bash
# Start the database
cd database
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## 🧪 Test Data

The database includes test users with password `password123`:

- `test@test.com`
- `test2@test.com`
- `test3@test.com`

## 📝 Usage Examples

### 1. User registration
```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"new@user.com","password":"my_password"}'
```

### 2. Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"new@user.com","password":"my_password"}'
```

### 3. Create post (using the login token)
```bash
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{"title":"My Post","content":"Post content"}'
```

## 🤝 Contributing

1. Fork the project
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

## 👨‍💻 Author

Developed by Erick Fariñas.
