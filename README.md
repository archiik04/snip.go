# snip.go: URL Shortener

A fast, minimal URL shortener built with **Go + Gin + Redis** on the backend and a clean vanilla HTML/CSS/JS frontend.

<img width="1189" height="690" alt="image" src="https://github.com/user-attachments/assets/eb85379a-d879-488c-aeff-dac048e93b9b" />


## Tech Stack

- **Backend:** Go, Gin, Redis (go-redis)
- **Frontend:** HTML, CSS, Vanilla JS
- **Backend Hosting:** Render
- **Frontend Hosting:** Vercel
- **Database:** Redis (Render Key Value)

## Features

-  Shorten any URL instantly
-  Custom aliases (e.g. `snip-go.onrender.com/my-link`)
-  Optional expiry (TTL in days)
-  Click tracking & stats
-  Copy to clipboard
-  Recent links history (saved in localStorage)

## Project Structure

```
snip-go/
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── handlers/
│   │   └── url.go
│   ├── store/
│   │   └── redis.go
│   └── utils/
│       └── shortcode.go
└── frontend/
    └── index.html
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/shorten` | Shorten a URL |
| `GET` | `/:code` | Redirect to original URL |
| `GET` | `/api/stats/:code` | Get click stats for a link |

### Shorten a URL

```bash
curl -X POST https://snip-go.onrender.com/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://google.com", "custom_code": "google", "ttl_days": 7}'
```

**Response:**
```json
{
  "code": "google",
  "short_url": "https://snip-go.onrender.com/google",
  "original_url": "https://google.com",
  "expires_in": 7
}
```

### Get Stats

```bash
curl https://snip-go.onrender.com/api/stats/google
```

**Response:**
```json
{
  "code": "google",
  "original_url": "https://google.com",
  "clicks": 42,
  "short_url": "https://snip-go.onrender.com/google"
}
```

## Running Locally

### Prerequisites

- Go 1.21+
- Docker 

### 1. Clone the repo

```bash
git clone https://github.com/himanshuraimau/snip-go.git
cd snip-go
```

### 2. Start Redis

```bash
docker run -d -p 6379:6379 redis:alpine
```

### 3. Run the backend

```bash
cd backend
cp .env.example .env  # edit if needed
go run main.go
```

### 4. Open the frontend

Open `frontend/index.html` directly in your browser.

The app will be running at `http://localhost:8080`.

## Environment Variables

Create a `.env` file:

```env
REDIS_URL=redis://localhost:6379
BASE_URL=http://localhost:8080
PORT=8080
```
