package main

import (
	"encoding/json"
	"f1/api"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static/css"))
	http.Handle("GET /static/css/", http.StripPrefix("/static/css/", fileServer))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		drivers, err := api.GetDriverData()
		if err != nil {
			http.Error(w, "Ошибка чтения шаблона"+err.Error(), http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(drivers)
		if err != nil {
			http.Error(w, "Ошибка JSON: "+err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Не удалось загрузить страницу", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, template.HTML(jsonBytes)); err != nil {
			log.Println("Ошибка рендеринга:", err)
		}
	})

	api.GetDriverData()

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("ошибка запуска", err)
	}
}
