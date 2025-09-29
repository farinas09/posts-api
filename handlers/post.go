package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/farinas09/rest-ws/middleware"
	"github.com/farinas09/rest-ws/models"
	"github.com/farinas09/rest-ws/repository"
	"github.com/farinas09/rest-ws/server"
)

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostResponse struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int64  `json:"user_id"`
}

func CreatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el UserId del contexto (ya validado por el middleware)
		userId, ok := middleware.GetUserIdFromContext(r.Context())
		if !ok {
			http.Error(w, "User ID not found in context", http.StatusInternalServerError)
			return
		}

		// Decodificar el request body
		var request = PostRequest{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Crear el post con el UserId del contexto
		post := models.Post{
			Title:   request.Title,
			Content: request.Content,
			UserId:  userId,
		}

		// Guardar en la base de datos
		if err := repository.CreatePost(r.Context(), &post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Responder con el post creado
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(PostResponse{
			Id:      post.Id,
			Title:   post.Title,
			Content: post.Content,
			UserId:  post.UserId,
		})
	}
}
