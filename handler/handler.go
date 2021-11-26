package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome to Home"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"),path.Join("views", "layouts.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]string {
	// 	"title" : "I'm Learning Golang Web",
	// 	"content" : "I'm Learning Golang Web",
	// }

		// data := entity.Product{ID:1, Name: "Mobilio", Price: 220000000, Stock: 3}

		// Slice
		data := []entity.Product{
			{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 11},
			{ID: 2, Name: "Xpander", Price: 270000000, Stock: 8},
			{ID: 3, Name: "Pajero Sport", Price: 50000000, Stock: 1},
		}

	tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello word, saya sedang belajar Golang web"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from Nintendo"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views","product.html"), path.Join("views", "layouts.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request){
	method := r.Method //GET, POST

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views","form.html"), path.Join("views","layouts.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		pesan := r.Form.Get("pesan")

		data := map[string]interface{}{
			"name" : name,
			"pesan" : pesan,
		}
		
		tmpl, err := template.ParseFiles(path.Join("views","result.html"), path.Join("views","layouts.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}  

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		return 
	}

	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}