package main

import (
	"net/http"

	"github.com/MdHasib01/go_social_app/internal/store"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `josn:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJOSON(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		//TODO: Changer after auth
		UserID: 1,
		Tags:   payload.Tags,
	}
	ctx := r.Context()
	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJOSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return

	}
}
