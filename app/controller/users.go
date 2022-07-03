package controller

import (
	"encoding/json"
	"github.com/EkaterinaShamanaeva/testWebProject/app/model"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

// GetUsers возвращает список всех пользователей
func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// получаем список всех пользователей
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	main := filepath.Join("public", "html", "usersDynamicPage.html")
	common := filepath.Join("public", "html", "common.html")
	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(rw, "users", users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func AddUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем значение из параметра name, переданного в форме запроса
	name := r.FormValue("Name")
	//получаем значение из параметра surname, переданного в форме запроса
	surname := r.FormValue("Surname")

	//проверяем на пустые значения
	if name == "" || surname == "" {
		http.Error(rw, "Имя и фамилия не могут быть пустыми", 400)
		return
	}
	//создаем новый объект
	user := model.NewUser(name, surname)
	//записываем нового пользователя в таблицу БД
	err := user.AddUser()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	err = json.NewEncoder(rw).Encode("Пользователь успешно добавлен!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
