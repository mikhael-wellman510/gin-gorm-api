Repository 
https://github.com/shayja/go-template-api/blob/main/cmd/app/app.go
https://github.com/Caknoooo/go-gin-clean-starter

install go gin 
go get -u github.com/gin-gonic/gin

install gorm dan driver sql
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

install dotenv
go get github.com/joho/godotenv

untuk menggunakan nilai desimal 
go get github.com/shopspring/decimal

untuk menggunakan UUID di Entity
go get github.com/google/uuid

Untuk JWT
go get -u github.com/golang-jwt/jwt/v5

untuk Swagger
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
Lalu tambahkan anotasi di setiap controlller

<!-- =================== -->
gunakan Panic untuk menghentikan aplikasi
- seperti config db 

gunakan Return kosong untuk menangani Error


<!-- Task -->

buat util , success response , errResponse , service , repository , crud

next -> Kirim gambar , Many To One(Relasi) ,  
