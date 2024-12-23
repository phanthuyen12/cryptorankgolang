package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// InitDB thiết lập kết nối PostgreSQL
func InitDB() {
	// Chuỗi kết nối PostgreSQL
	dsn := "postgres://username:password@localhost:5432/dbname" // Thay đổi `username`, `password` và `dbname` cho phù hợp
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Không thể parse connection string: %v\n", err)
	}

	// Tạo kết nối pool
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Không thể kết nối đến cơ sở dữ liệu: %v\n", err)
	}

	log.Println("Kết nối đến PostgreSQL thành công!")
}

// CloseDB đóng kết nối database khi chương trình kết thúc
func CloseDB() {
	DB.Close()
}
