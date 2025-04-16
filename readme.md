# URL Shortener Service with Analytics 🚀 

![Go](https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)

## 📋 Overview

A high-performance URL shortening service built with Go that not only creates shortened URLs but also provides robust analytics for tracking link usage.

### ✨ Features

- ⚡ Fast URL shortening with custom or auto-generated aliases
- 📊 Comprehensive analytics for link clicks
- 🔄 Redirect handling with proper HTTP status codes
- 🔍 URL validation and sanitization
- 🧪 Thoroughly tested codebase

## 🚀 Getting Started

### Prerequisites

- Go 1.16+
- Internet connection (for dependency downloads)

### Installation

1. Clone the repository
```
git clone https://github.com/yourusername/url-shortener-analytics.git
cd url-shortener-analytics
```

2. Install dependencies
```
go mod download
```

3. Run the server
```
go run main.go
```

## 🔧 Usage

### Creating a Short URL

```
POST /shorten
```

Request body:
```
{
    "url": "https://example.com/very-long-url-that-needs-shortening",
    "custom_alias": "mylink"  // Optional
}
```

Response:
```
{
    "short_url": "http://yourdomain.com/mylink",
    "original_url": "https://example.com/very-long-url-that-needs-shortening",
    "created_at": "2023-04-12T15:04:05Z"
}
```

### ✅ Health Check

The service provides a health endpoint to verify it's running properly:

```
GET /health
```

## 🧪 Testing

Run tests with:

```
go test ./...
```

## 📈 Future Improvements

- [ ] User authentication
- [ ] Rate limiting
- [ ] Expiry for short URLs
- [ ] Dashboard for analytics visualization
- [ ] API key management

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👨‍💻 Author

Made with ❤️ by KunjShah95