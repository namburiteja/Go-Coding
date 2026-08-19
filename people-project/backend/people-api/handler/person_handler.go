package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"people-api/dto"
	"strconv"
	"strings"

	"people-api/service"
)

type PersonHandler struct {
	Service *service.PersonService
}

func NewPersonHandler(svc *service.PersonService) *PersonHandler {
	return &PersonHandler{
		Service: svc,
	}
}

type PaginatedResponse struct {
	Data       []dto.PersonDTO `json:"data"`
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
	Total      int64           `json:"total"`
	TotalPages int             `json:"totalPages"`
}

func (h *PersonHandler) GetAllPeople(w http.ResponseWriter, r *http.Request) {

	page := 1
	limit := 20

	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	var err error

	if pageParam != "" {
		page, err = strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			http.Error(w, "Invalid page", http.StatusBadRequest)
			return
		}
	}

	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam)
		if err != nil || limit < 1 || limit > 100 {
			http.Error(w, "Limit must be between 1 and 100", http.StatusBadRequest)
			return
		}
	}

	offset := (page - 1) * limit

	people, err := h.Service.GetPeoplePaginated(
		r.Context(),
		int32(limit),
		int32(offset),
	)

	if err != nil {
		log.Println("GetPeoplePaginated error:", err)
		http.Error(w, "Failed to fetch people", http.StatusInternalServerError)
		return
	}

	total, err := h.Service.CountPeople(r.Context())

	if err != nil {
		log.Println("CountPeople error:", err)
		http.Error(w, "Failed to count people", http.StatusInternalServerError)
		return
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))
	if page > totalPages && totalPages > 0 {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	response := PaginatedResponse{
		Data:       people,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func (h *PersonHandler) GetPeopleByID(w http.ResponseWriter, r *http.Request) {
	playerID := strings.TrimPrefix(r.URL.Path, "/people/")

	if playerID == "" {
		http.Error(w, "Player ID is required", http.StatusBadRequest)
		return
	}

	person, err := h.Service.GetPeopleByID(r.Context(), playerID)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Person not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to fetch person", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func (h *PersonHandler) SearchPeopleByName(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := strings.TrimSpace(
		r.URL.Query().Get("name"),
	)

	if len(name) < 2 {
		http.Error(
			w,
			"Search requires at least 2 characters",
			http.StatusBadRequest,
		)
		return
	}

	people, err := h.Service.SearchPeopleByName(
		r.Context(),
		name,
	)

	if err != nil {
		log.Println("SearchPeopleByName error:", err)
		http.Error(
			w,
			"Failed to search people",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(people)
}
