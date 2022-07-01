package main

import (
	"github.com/EkaterinaShamanaeva/testWebProject/app/controller"
	"github.com/EkaterinaShamanaeva/testWebProject/app/server"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	//инициализируем подключение к базе данных
	err := server.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	// создаем роутер для обслуживания запросов
	router := httprouter.New()
	routes(router)

	// прикрепляемся к хосту и свободному порту
	err = http.ListenAndServe("localhost:4444", router)
	if err != nil {
		log.Fatal(err)
	}

}
func routes(router *httprouter.Router) {
	// путь к папке с html, css и др. файлами
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	// что следует выполнять при входящих запросах по указанному адресу
	router.GET("/", controller.StartPage)
	router.GET("/users", controller.GetUsers)
}
