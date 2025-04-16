# Generic HTTP Server with Blog Post Functionality 🚀

![Go](https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)

## 📋 Overview

A versatile HTTP server built with Go that provides blog post functionality with clean architecture, easy customization, and performance optimization.

### ✨ Features

- ⚡ Fast and lightweight HTTP server implementation
- 📝 Complete blog post CRUD operations
- 🔍 Search and filtering capabilities
- 🔒 Content validation and sanitization
- 🧪 Thoroughly tested codebase

## 🚀 Getting Started

### Prerequisites

- Go 1.16+
- Internet connection (for dependency downloads)

### Installation

1. Clone the repository
```
git clone https://github.com/KunjShah95/generic-http-blog-server.git
cd generic-http-blog-server
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

### Creating a Blog Post

```
POST /posts
```

Request body:
```
{
    "title": "My First Blog Post",
    "content": "This is the content of my first blog post.",
    "tags": ["golang", "web-dev"]
}
```

Response:
```
{
    "id": "123",
    "title": "My First Blog Post",
    "content": "This is the content of my first blog post.",
    "tags": ["golang", "web-dev"],
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
- [ ] Comments functionality
- [ ] Media uploads
- [ ] Admin dashboard
- [ ] API key management

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👨‍💻 Author

Made with ❤️ by KunjShah95