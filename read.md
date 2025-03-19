Repository 
https://github.com/shayja/go-template-api/blob/main/cmd/app/app.go


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


<!-- =================== -->
gunakan Panic untuk menghentikan aplikasi
- seperti config db 

gunakan Return kosong untuk menangani Error


<!-- Task -->

buat util , success response , errResponse , service , repository , crud
