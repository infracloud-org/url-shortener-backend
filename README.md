# URL Shortener Service

A simple URL shortener service built using **Go (Echo Framework)** for the backend and **React.js** for the frontend.  
The service allows users to shorten URLs, track analytics, and retrieve original URLs.

## Features

- Shorten long URLs to compact short URLs.
- Redirect short URLs to their original destination.
- Track analytics (top 3 most shortened domains).
- In-memory storage using `sync.Map` for fast performance.
- RESTful API built with the Echo framework.
- Containerized using Docker for easy deployment.

## Tech Stack
### Backend
- **Language**: Go (Echo Framework)
- **Storage**: In-memory (`sync.Map`)
- **Libraries**: Echo, Base62 Encoding, URL Parsing
- **Testing**: Go Unit Tests (`testing` package)
- **Deployment**: Docker (Extensible to Kubernetes)

### Frontend
- **Framework**: React.js
- **State Management**: Redux (Optional)
- **UI Components**: Material UI

## Installation and Setup
### Prerequisites
Before begin with the **URL Shortening Service**, ensure the following:
- Install **Go** (>= 1.20)
- Install **Docker**
- Install **Postman** (for API testing)

### Getting Started
1. **Clone the repository**:
    ```bash
    git clone https://github.com/infracloud-org/url-shortener-backend.git
    cd url-shortener-backend
    ```
2. **Install Dependencies**:
    ```bash
    go mod tidy
    ```
3. **Run the URL Shortener Backend Service**:
- If using docker:
    ```bash
    docker build -t url-shortener-backend:1.0.0 -f ./Dockerfile .

    docker run -itd --name url-shortener-backend -p 9090:9090 url-shortener-backend:1.0.0
    ```
- If running locally:
    ```bash
    go run cmd/main.go
    ```

## API Endpoints
1. Shorten a URL
    ```http
    curl -v -X POST http://localhost:9090/api/v1/shorten -d '{ "originalURL": "https://example.com" }'
    ```
    Response:
    ```json
    {
        "shortURL": "ij4HVY"
    }
    ```
2. Redirect to Original URL
    ```http
    curl -v -X GET http://localhost:9090/api/v1/:shortCode
    ```
    Response:
    Redirects to https://example.com
3. Get Top 3 Shortened Domains
    ```http
    curl -v -X GET http://localhost:9090/api/v1/metrics
    ```
    Response:
    ```json
    {
        "Udemy": 6,
        "YouTube": 4,
        "Wikipedia": 2
    }
    ```

## License
This project is licensed under **Apache 2.0 License** - see LICENSE.md file for details