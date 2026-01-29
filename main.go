package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"tugas-1/database"
	"tugas-1/handlers"
	"tugas-1/models"
	"tugas-1/repositories"
	"tugas-1/services"

	"github.com/spf13/viper"
)

var kategori = []models.Category {
	{ID: 1, Name: "Action", Description: "Film yang melibatkan banyak aksi atau combat yang spektakuler"},
	{ID: 2, Name: "Horror", Description: "Film yang membuat bulu kuduk merinding dan menyerang psikologis"},
}

type Config struct {
	Port string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}


func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port: viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// setup db
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database", err)
	}
	defer db.Close()
	// root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Welcome to Category API",
		})
	}) 

	// dependency injection category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Setup routes category
	http.HandleFunc("/api/category", categoryHandler.HandleCategories)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)

	// dependency injection product 
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Setup routes product
	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	fmt.Println("Server running di localhost:" + config.Port)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}

