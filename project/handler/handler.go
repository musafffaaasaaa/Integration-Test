package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/musa/project/db"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product

	// Распарсить JSON-данные запроса в структуру product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получить клиент MongoDB
	client, err := db.ConnectToMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	// Вызвать функцию для создания продукта
	err = db.InsertProduct(client, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	// Add other fields as needed
}

var products []Product

func init() {

	products = append(products, Product{ID: 1, Name: "Product 1", Price: 19.99})
	products = append(products, Product{ID: 2, Name: "Product 2", Price: 29.99})
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {

	productID := 1
	product := findProductByID(productID)

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func findProductByID(id int) *Product {
	for _, p := range products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
