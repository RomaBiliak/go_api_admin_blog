# GO REST API admin blog

## Dependency

* github.com/gorilla/mux
* github.com/dgrijalva/jwt-go
* golang.org/x/crypto/bcrypt
* github.com/mattn/go-sqlite
* github.com/joho/godotenv
* github.com/disintegration/imaging

## Directory structure
```
go_api_admin_blog
│   .env
│   main.go
│   README.md
│
├───api
│   ├───auth
│   │       token.go
│   │
│   ├───controllers
│   │       file_controller.go
│   │       filter_controller.go
│   │       image_controller.go
│   │       login_controller.go
│   │       routes.go
│   │       server.go
│   │       user_controller.go
│   │
│   ├───fille
│   ├───middlewares
│   │       middlewares.go
│   │
│   ├───models
│   │       Filter.go
│   │       Image.go
│   │       User.go
│   │
│   └───responses
│           json.go
│
├───bootstrap
│       run.go
│
├───db
│       database.db
│
├───image
├───library
│   ├───file
│   │       file.go
│   │
│   └───image
│           image.go
│
├───server
│       server.go
│
├───upload_photo
└───watermark
        watermark.png
```