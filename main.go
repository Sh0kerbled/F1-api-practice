package main

import (
	"f1/api"
)

func main() {
	api.GetDriverData()

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	drivers, err := api.GetDriver()
	// 	if err != nil {
	// 		http.Error(w, "Ошибка чтения шаблона"+err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	tmpl, err := template.ParseFiles("templates/index.html")
	// 	if err != nil {
	// 		http.Error(w, "Не удалось загрузить страницу", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	if err := tmpl.Execute(w, drivers); err != nil {
	// 		log.Println("Ошибка рендеринга:", err)
	// 	}
	// })

	// api.GetDriverData()

	// if err := http.ListenAndServe(":8000", nil); err != nil {
	// 	fmt.Println("ошибка запуска", err)
	// }
}
