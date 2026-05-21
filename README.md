
# snip.go
<div align="center">
    A fast, minimal URL shortener: paste a link, get one that doesn't overstay its welcome.
    

Built with Go + Gin on the backend, Redis for storage, and a clean vanilla HTML/CSS/JS frontend.

<img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white" />
<img src="https://img.shields.io/badge/Gin-Framework-00A9A5?style=flat-square" />
<img src="https://img.shields.io/badge/Redis-Key--Value-DC382D?style=flat-square&logo=redis&logoColor=white" />
<img src="https://img.shields.io/badge/Hosted_on-Render-46E3B7?style=flat-square&logo=render&logoColor=white" />
<img src="https://img.shields.io/badge/Frontend-Vercel-000000?style=flat-square&logo=vercel&logoColor=white" />


<br />

</div>

## Interface

<div align="center">
  <img src="https://github.com/user-attachments/assets/ba2e5871-bab0-417f-9a37-9487c09dca8d" width="520" alt="snip.go — main interface" />
  <br /><sub>Main interface</sub>
  <br /><br />
  <img src="https://github.com/user-attachments/assets/264fa76f-5fcc-4e92-a5d4-aa4d7ab0273f" width="520" alt="snip.go — after shortening" />
  <br /><sub>After shortening — with copy, stats, QR code, and history</sub>
</div>

<br />


## Features

- **Instant shortening** — paste any URL, get a clean short link in one click
- **Custom aliases** — choose your own slug (e.g. `snip-go.onrender.com/my-link`)
- **Expiry / TTL** — optionally expire a link after N days
- **Click tracking** — per-link stats with a click counter
- **QR code** — generate and download a QR code for any shortened link
- **Clipboard detect** — auto-fills the input if your clipboard holds a URL
- **Recent history** — last 20 links saved in `localStorage`, no login needed


## Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go 1.21+ |
| HTTP Framework | Gin |
| Database | Redis (via go-redis) |
| Backend Hosting | Render |
| Frontend Hosting | Vercel |
| Frontend | Vanilla HTML / CSS / JS |


## Project Structure

```
snip-go/
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── handlers/
│   │   └── url.go          # shorten, redirect, stats handlers
│   ├── store/
│   │   └── redis.go        # Redis client + operations
│   └── utils/
│       └── shortcode.go    # random code generation
└── frontend/
    └── index.html          # single-file frontend
```


## API Reference

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/shorten` | Shorten a URL |
| `GET` | `/:code` | Redirect to the original URL |
| `GET` | `/api/stats/:code` | Get click stats for a link |


### `POST /api/shorten`

**Request**

```bash
curl -X POST https://snip-go.onrender.com/api/shorten \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://google.com",
    "custom_code": "google",
    "ttl_days": 7
  }'
```

All fields except `url` are optional.

| Field | Type | Description |
|-------|------|-------------|
| `url` | `string` | **Required.** The URL to shorten |
| `custom_code` | `string` | Custom slug (auto-generated if omitted) |
| `ttl_days` | `int` | Days until expiry (no expiry if omitted) |

**Response**

```json
{
  "code": "google",
  "short_url": "https://snip-go.onrender.com/google",
  "original_url": "https://google.com",
  "expires_in": 7
}
```

---

### `GET /api/stats/:code`

```bash
curl https://snip-go.onrender.com/api/stats/google
```

**Response**

```json
{
  "code": "google",
  "original_url": "https://google.com",
  "clicks": 42,
  "short_url": "https://snip-go.onrender.com/google"
}
```

---

## Running Locally

### Prerequisites

- Go 1.21+
- Docker (for Redis)

### 1. Clone the repo

```bash
git clone https://github.com/himanshuraimau/snip-go.git
cd snip-go
```

### 2. Start Redis

```bash
docker run -d -p 6379:6379 redis:alpine
```

### 3. Configure environment

```bash
cd backend
cp .env.example .env
```

```env
REDIS_URL=redis://localhost:6379
BASE_URL=http://localhost:8080
PORT=8080
```

### 4. Run the backend

```bash
go run main.go
```

### 5. Open the frontend

Open `frontend/index.html` directly in your browser (no build step needed)

The API will be running at `http://localhost:8080`.

