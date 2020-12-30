package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	UID         string   `json: "UID"`
	Title       string   `json: "Title"`
	Description string   `json: Description`
	Thumbnail   string   `json: "Thumbnail"`
	Article     Category `json : "article"`
}

type Category struct {
	Name         string `json: "Name"`
	Code         int    `json: "code"`
	Availability bool   `json: "Availability"`
}

var articles []Article = []Article{}

func addArticle(w http.ResponseWriter, r *http.Request) {

	var newArticle Article
	json.NewDecoder(r.Body).Decode(&newArticle)

	articles = append(articles, newArticle)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(articles)

}

func getAllArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(articles) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	article := articles[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)

}

func updateArticle(w http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(articles) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	var updatedArticle Article

	json.NewDecoder(r.Body).Decode(&updatedArticle)

	articles[id] = updatedArticle

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedArticle)
}

func patchArticle(w http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(articles) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	article := &articles[id]
	json.NewDecoder(r.Body).Decode(article)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return

	}
	if id >= len(articles) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	articles = append(articles[:id], articles[id+1:]...)

	w.WriteHeader(200)

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/article", addArticle).Methods("POST")

	router.HandleFunc("/article", getAllArticle).Methods("GET")

	router.HandleFunc("/article/{id}", getArticle).Methods("GET")

	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

	router.HandleFunc("/article/{id}", patchArticle).Methods("PATCH")

	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	http.ListenAndServe(":5000", router)

}
