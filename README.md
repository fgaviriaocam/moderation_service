*IMPORTANT*
This basic proyect is for show ddd (https://www.paradigmadigital.com/dev/ddd-dominio-implica-crecer-fuerte/) with golang, remember, ddd can be implemented in many ways depending on the author!

# Web Service Moderation
WS for make moderation in text

# Start 🚀
    1. Clone this project -> https://github.com/fgaviriaocam/moderation_service
    2. Make sure port 8080 is not busy.
    3. go get all 
    4. go run main.go

# Pre-requirements 📋
It is necessary to install -> https://golang.org/ 

# Dependencies 🤝
**IMPORTAT:** All dependencies will register and will be usable all dependency, however all dependency that is not used and is imported they should delete from de mod file

- [github.com/gin-gonic/gin] Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [github.com/gin-contrib/pprof] Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.
- [github.com/swaggo/gin-swagger] [github.com/swaggo/files] This is a open api for documentation a some tests 

# Project Structure 🧱
```
moderation_service
├─ .env
├─ Dockerfile
├─ README.md
├─ application
│  └─ ModerationController.go
├─ constants
│  └─ constans.go
├─ domain
│  ├─ models
│  │  ├─ Language.go
│  │  ├─ Moderation.go
│  │  └─ Response.go
│  └─ service
│     ├─ ModerationService.go
│     ├─ ModerationServiceImpl.go
├─ go.mod
├─ go.sum
├─ infrastructure
│  ├─ persistence
│  │  ├─ DbHelper.go
│  │  └─ ModerationRepositoryImpl.go
│  └─ repository
│     └─ ModerationRepository.go
├─ interfaces
│  └─ middleware
│     ├─ CORSMiddleware.go
│     └─ server
│        ├─ Server.go
│        └─ ServerImpl.go
├─ languages
│  ├─ en.json
│  ├─ es.json
│  └─ pt.json
├─ main.go
└─ utils
   └─ converter.go
```

# Built with 🛠️
    - Visual Studio Code
    - Goland

# Endpoints
    - POST /moderation

# Authors
Andrés Gaviria
Software Engineer