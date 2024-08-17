package connection

import (
	"database/sql"
	"fmt"
	"golang_biomtrik_login_fido/internal/config"

	"github.com/gofiber/fiber/v2/log"

	// kita import juga library postgresql untuk bisa connect
	_ "github.com/lib/pq"
)

func GetDatabase(conf *config.Databases) *sql.DB {
	// kalau s print ini untuk membuat return nilai
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		conf.Host, conf.User, conf.Password, conf.Name, conf.Port)
	db, err := sql.Open("postgres", dsn) // disini kita pakai postgres, yg mana datasourcename nya dari dsn yg sudah kita buat
	if err != nil {
		log.Fatal("failed to open connection", err.Error())
	}

	err = db.Ping() // kita cek koneksi ke db
	if err != nil {
		log.Fatal("failed to ping database", err.Error())
	}
	return db
}
