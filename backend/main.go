package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    posts,
	})
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	post.ID = fmt.Sprintf("%d", len(posts)+1)
	posts = append(posts, post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Post created successfully",
		Data:    post,
	})
}

func getPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, post := range posts {
		if post.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{
				Success: true,
				Data:    post,
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: "Post not found",
	})
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, post := range posts {
		if post.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{
				Success: true,
				Message: "Post deleted successfully",
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: "Post not found",
	})
}

func main() {
	posts = append(posts, Post{ID: "1", Title: "First Post", Content: "This is the first post!", Author: "John Doe"})

	r := mux.NewRouter()

	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := corsHandler.Handler(r)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
