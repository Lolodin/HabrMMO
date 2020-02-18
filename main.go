package main

import (
	"exampleMMO/GameController"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// Роутер для доступа к клиенту
	http.HandleFunc("/map", GameController.Map_Handler)
	http.HandleFunc("/", indexHandler)
	// Открываем доступ к статичным ресурсам (скрипты, картинки и тд.)
	http.Handle("/node_modules/phaser/dist/", http.StripPrefix("/node_modules/phaser/dist/", http.FileServer(http.Dir("./node_modules/phaser/dist/"))))
	http.Handle("/Client/", http.StripPrefix("/Client/", http.FileServer(http.Dir("./Client/"))))
	http.Handle("/Client/Content/", http.StripPrefix("/Client/Content/", http.FileServer(http.Dir("./Client/Content/"))))
	// Запускаем сервер. Указываем любой порт, я выбрал 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Обработчик для index.html, здесь мы просто отдаем клиент пользователю
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indexAction")
	t, _ := template.ParseFiles("index.html")
	err := t.Execute(w, "index")
	if err != nil {
		fmt.Println(err.Error())
	}
}
