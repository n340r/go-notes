package httpserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/n340r/backend-notes/toptal/internal/app/common/server"
	"github.com/n340r/backend-notes/toptal/internal/app/domain"
)

func (h HttpServer) GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	categoryID, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		server.BadRequest("invalid-category-id", err, w, r)
		return
	}
	category, err := h.categoryService.GetCategory(r.Context(), categoryID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.BadRequest("category-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseCategory(category)

	server.RespondOK(response, w, r)
}

func (h HttpServer) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryRequest CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}

	if err := categoryRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	category, err := domain.NewCategory(domain.NewCategoryData{
		Name: categoryRequest.Name,
	})
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	insertedCategory, err := h.categoryService.CreateCategory(r.Context(), category)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseCategory(insertedCategory)

	server.RespondOK(response, w, r)
}

func (h HttpServer) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// get vars
	// get categoryId from somwhere
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		server.BadRequest("invalid-category-id", err, w, r)
		return
	}

	// decode
	var categoryRequest CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}
	// validate
	if err := categoryRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	// check category exists
	_, err = h.categoryService.GetCategory(r.Context(), categoryID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.BadRequest("category-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	// format ncoming category into domain format
	category, err := domain.NewCategory(domain.NewCategoryData{
		ID:   categoryID,
		Name: categoryRequest.Name,
	})
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	// update via domain.UpdateCategory something like this - pass context, category
	updatedCategory, err := h.categoryService.UpdateCategory(r.Context(), category)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	// serialize response
	// respondOK
	response := toResponseCategory(updatedCategory)

	server.RespondOK(response, w, r)
}

func (h HttpServer) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	categoryID, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		server.BadRequest("invalid-category-id", err, w, r)
		return
	}

	_, err = h.categoryService.GetCategory(r.Context(), categoryID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("category-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	err = h.categoryService.DeleteCategory(r.Context(), categoryID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("category-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	server.RespondOK(map[string]bool{"deleted": true}, w, r)
}

func (h HttpServer) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.GetCategories(r.Context())
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := make([]CategoryResponse, 0, len(categories))
	for _, category := range categories {
		response = append(response, toResponseCategory(category))
	}

	server.RespondOK(response, w, r)
}
