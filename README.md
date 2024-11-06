# KeepActive Backend

KeepActive is a service that helps keep free-tier hosted websites active by sending periodic GET requests. This prevents them from going dormant due to inactivity.

## Features

- User authentication with JWT
- Website monitoring management
- Automated GET request scheduling
- RESTful API design
- Docker containerization
- Production-ready configuration

## Tech Stack

- Go 1.23
- Gin Web Framework
- GORM with PostgreSQL
- JWT Authentication
- Docker
- Supabase (Database)
- Render (Backend Deployment)

## Local Development Setup

### Prerequisites

- Go 1.23 or higher
- Docker
- PostgreSQL or Supabase account
- Git

### Installation

1. Clone the repository

```bash
git clone https://github.com/oadultradeepfield/keepactive-backend.git
cd keepactive-backend
```

2. Install dependencies

```bash
go mod tidy
```

3. Create `.env` file in the root directory

```env
DATABASE_URL=your_supabase_postgres_connection_string
JWT_SECRET=your_secure_jwt_secret
ALLOWED_ORIGINS=your_fronted_url
GO_ENV=development
PORT=8080
```

4. Run the application

```bash
go run main.go
```

### Docker Local Development

1. Build the Docker image

```bash
docker build -t keepactive-backend .
```

2. Run the container

```bash
docker run -p 8080:8080 --env-file .env keepactive-backend
```

## API Endpoints

### Authentication Endpoints

- `POST /api/register` - Register new user

  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```

- `POST /api/login` - Login user
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```

### Website Management Endpoints

All these endpoints require JWT Authentication header: `Authorization: Bearer <token>`

- `POST /api/websites` - Create new website

  ```json
  {
    "name": "My Website",
    "url": "https://mywebsite.com",
    "duration": 7
  }
  ```

- `GET /api/websites` - List all websites

- `DELETE /api/websites/:id` - Delete website

## Deployment to Render

### Prerequisites

1. Create a [Render](https://render.com) account
2. Create a [Supabase](https://supabase.com) account and database
3. Have your code pushed to GitHub

### Deployment Steps

1. **Create New Web Service**

   - Go to Render Dashboard
   - Click "New +" and select "Web Service"
   - Connect your GitHub repository

2. **Configure Web Service**

   - Name: `keepactive-backend` (or your preferred name)
   - Region: Choose nearest to your target audience
   - Runtime: Docker
   - Instance Type: Free
   - Health Check Path: `/health`

3. **Environment Variables**
   Add the following environment variables in Render dashboard:

   ```
   DATABASE_URL=your_supabase_postgres_connection_string
   JWT_SECRET=your_secure_jwt_secret
   ALLOWED_ORIGINS=https://your-frontend-domain.com
   GO_ENV=production
   PORT=8080
   ```

4. **Deploy**
   - Click "Create Web Service"
   - Render will automatically build and deploy your application

### Checking Deployment

1. Your app will be available at `https://your-service-name.onrender.com`
2. Test the health check endpoint: `https://your-service-name.onrender.com/health`
3. Try registering a user using the `/api/register` endpoint

## Project Structure

```
keepactive-backend/
├── config/
│   ├── config.go    # Configuration management
│   └── database.go  # Database initialization
├── handlers/
│   ├── auth.go      # Authentication handlers
│   └── website.go   # Website management handlers
├── middleware/
│   ├── auth.go      # JWT authentication middleware
│   └── cors.go      # CORS middleware
├── models/
│   └── models.go    # Data models
├── services/
│   └── pinger.go    # Website pinger service
├── main.go          # Application entry point
├── Dockerfile       # Docker configuration
├── go.mod          # Go modules file
├── go.sum          # Go modules checksum
└── .env            # Environment variables (local only)
```
