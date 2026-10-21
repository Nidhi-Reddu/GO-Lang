package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"temp/database"
	"temp/models"

	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Unable to connect DataBase", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `SELECT id, title, content, author, createdAt, updatedAt FROM public.posts ORDER BY id DESC`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// var posts []models.Post
	posts := make([]models.Post, 0)
	for rows.Next() {
		var post models.Post

		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.Created_at, &post.Updated_at)
		if err != nil {
			http.Error(w, "Failed to read posts", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error while reading posts", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(posts)
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post_id := vars["id"]
	// post_id := r.URL.Query().Get("id")
	// post_id := r.PathValue("id")
	if post_id == "" {
		http.Error(w, "Post does not exist !", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(post_id, 10, 64)
	if err != nil {
		http.Error(w, "Post does not exist !", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Unable to connect DataBase", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var post models.Post
	query := `SELECT id, title, content, author FROM public.posts WHERE id = $1`
	err = db.QueryRow(query, id).Scan(&post.Id, &post.Title, &post.Content, &post.Author)

	if err != nil {
		http.Error(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO posts (title, content, author) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err = db.QueryRow(query, post.Title, post.Content, post.Author).Scan(&id)
	if err != nil {
		log.Printf("CreatePost error: %v", err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Post created successfully",
		"id":      id,
	})

}
func DeletePosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post_id := vars["id"]
	// post_id := r.URL.Query().Get("id")
	if post_id == "" {
		http.Error(w, "Post not found", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(post_id, 10, 64)
	if err != nil {
		http.Error(w, "Post does not exist", http.StatusNotFound)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	query := `DELETE FROM public.posts WHERE id = $1`
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post_id := vars["id"]

	if post_id == "" {
		http.Error(w, "Post does not exist !", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(post_id, 10, 64)
	if err != nil {
		http.Error(w, "Post does not exist !", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Unable to connect DataBase", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	query := `UPDATE public.posts SET title = $1, content = $2, author = $3, updatedAt = CURRENT_TIMESTAMP WHERE id = $4`
	_, err = db.Exec(query, post.Title, post.Content, post.Author, id)
	if err != nil {
		log.Printf("UpdatePost error: %v", err)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Post Updated successfully",
	})
}
