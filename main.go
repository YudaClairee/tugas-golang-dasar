package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

var kategori = []Category {
	{ID: 1, Name: "Action", Description: "Film yang melibatkan banyak aksi atau combat yang spektakuler"},
	{ID: 2, Name: "Horror", Description: "Film yang membuat bulu kuduk merinding dan menyerang psikologis"},
}



func main() {
	// root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Welcome to API",
		})
	}) 

	// /api/category
	http.HandleFunc("/api/category", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(kategori)
		} else if r.Method == "POST" {
		var kategoriBaru Category
		err := json.NewDecoder(r.Body).Decode(&kategoriBaru)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		kategoriBaru.ID = len(kategori) + 1
		kategori = append(kategori, kategoriBaru)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
		json.NewEncoder(w).Encode(kategoriBaru)
	}
	})


	// /api/category/{id}
	http.HandleFunc("/api/category/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryById(w, r)
		} else if r.Method == "PUT" {
			updateCategoryById(w, r)
		} else if r.Method == "DELETE" {
			deleteCategoryById(w, r)
		}
	})

	fmt.Println("Server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}

func getCategoryById (w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid category id", http.StatusBadRequest)
			return 
		}
		for _, k := range kategori {
			if k.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(k)
				return
			}
		}
		http.Error(w, "kategori tidak ditemukan", http.StatusNotFound)
	
}

func deleteCategoryById (w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid category id", http.StatusBadRequest)
			return
		}

		for i, k := range kategori {
			if k.ID == id {
				kategori = append(kategori[:i], kategori[i+1:]... )

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"message": "kategori berhasil dihapus",
				})
				return
			}
		}
		http.Error(w, "kategori belum ada", http.StatusNotFound)
}

func updateCategoryById (w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid category id", http.StatusBadRequest)
		}

		// get data dari request
		var updateKategori Category
		err = json.NewDecoder(r.Body).Decode(&updateKategori)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		for i := range kategori {
			if kategori[i].ID == id {
				updateKategori.ID = id
				kategori[i] = updateKategori
	
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(updateKategori)
				return
			}
		}
		http.Error(w, "kategori belum ada", http.StatusNotFound)
}