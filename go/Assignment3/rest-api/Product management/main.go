import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	UID         string  `json :"UID"`
	ProductName string  `json :"ProductName"`
	Description string  `json :"Description"`
	Price       float64 `json :"Price"`
}

var storage []Product

//homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welocme to the Home Page: homePage()")
}

//getproduct
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Welcome to Get Product : getProduct()")

	json.NewEncoder(w).Encode(storage)
}

//createproduct
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	storage = append(storage, product)

	json.NewEncoder(w).Encode(product)

}

//deleteProduct
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteProductAtUid(params["uid"])

	json.NewEncoder(w).Encode(storage)

}

//deleteAtUID
func _deleteProductAtUid(uid string) {

	for index, product := range storage {
		if product.UID == uid {
			storage = append(storage[:index], storage[index+1:]...)
			break

		}
	}
}

//updateProduct
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	params := mux.Vars(r)

	_deleteProductAtUid(params["uid"])
	storage = append(storage, product)
	json.NewEncoder(w).Encode(storage)

}

//handleRequests
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/product_storage", getProduct).Methods("GET")
	router.HandleFunc("/product_storage", createProduct).Methods("POST")
	router.HandleFunc("/product_storage/{uid}", deleteProduct).Methods("DELETE")
	router.HandleFunc("/product_storage/{uid}", updateProduct).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}

//main function
func main() {
	storage = append(storage, Product{
		UID:         "0",
		ProductName: "Buiscuit",
		Description: "Hand tossed buiscuit",
		Price:       20.50,
	})

	handleRequests()
