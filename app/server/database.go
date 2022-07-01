package server

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = ""
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test_web_go"
)

// Db глобальная переменная с подключением к БД
var Db *sql.DB

// InitDB функция, инициирующая подключение к БД
func InitDB() (err error) {
	//строка, содержащая данные для подключения к БД в следующем формате:
	//login:password@tcp(host:port)/dbname
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/",
		dbName, "?charset=utf8"}, "")

	// Открываем базу данных, первое - имя драйвера, поэтому import: _ "github.com/go-sql-driver/mysql"
	Db, _ = sql.Open("mysql", path)
	// Устанавливаем максимальное количество подключений к базе данных
	Db.SetConnMaxLifetime(100)
	// Устанавливаем максимальное количество неактивных подключений к базе данных
	Db.SetMaxIdleConns(10)
	// Проверяем соединение
	if err := Db.Ping(); err != nil {
		fmt.Println("open database fail")
		return err
	}
	fmt.Println("connect success")
	return nil
}
