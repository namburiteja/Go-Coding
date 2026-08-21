package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"people-api/db/generated"
	"people-api/dto"
	"people-api/service"
	"people-api/token"
	"strconv"
	"strings"
	"time"
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

type CursorPaginatedResponse struct {
	Data       []dto.PersonDTO `json:"data"`
	NextCursor string          `json:"nextCursor,omitempty"`
	HasNext    bool            `json:"hasNext"`
}

type TokenPaginatedResponse struct {
	Data      []dto.PersonDTO `json:"data"`
	NextToken string          `json:"nextToken,omitempty"`
	HasNext   bool            `json:"hasNext"`
}


func (h *PersonHandler) GetPeopleByID(w http.ResponseWriter, r *http.Request) {

	playerID := strings.TrimPrefix(r.URL.Path, "/people/")

	if playerID == "" {
		http.Error(w, "Player ID is required", http.StatusBadRequest)
		return
	}

	person, err := h.Service.GetPeopleByID(
		r.Context(),
		playerID,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			http.Error(
				w,
				"Person not found",
				http.StatusNotFound,
			)
			return
		}

		http.Error(
			w,
			"Failed to fetch person",
			http.StatusInternalServerError,
		)
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

func (h *PersonHandler) GetPeopleCursor(
	w http.ResponseWriter,
	r *http.Request,
) {

	limit := 20

	limitParam := r.URL.Query().Get("limit")

	if limitParam != "" {

		parsedLimit, err := strconv.Atoi(limitParam)

		if err != nil || parsedLimit < 1 || parsedLimit > 100 {
			http.Error(
				w,
				"Limit must be between 1 and 100",
				http.StatusBadRequest,
			)
			return
		}

		limit = parsedLimit
	}

	cursor := r.URL.Query().Get("cursor")

	/*
		Ask for one extra record.

		Example:
		limit = 10

		database returns 11

		10 → data
		1  → proves another batch exists
	*/

	people, err := h.Service.GetPeopleCursor(
		r.Context(),
		cursor,
		int32(limit+1),
	)

	if err != nil {
		log.Println("GetPeopleCursor error:", err)

		http.Error(
			w,
			"Failed to fetch people",
			http.StatusInternalServerError,
		)
		return
	}

	hasNext := len(people) > limit

	if hasNext {
		people = people[:limit]
	}

	response := CursorPaginatedResponse{
		Data:    people,
		HasNext: hasNext,
	}

	if hasNext {
		lastPerson := people[len(people)-1]

		response.NextCursor = lastPerson.PlayerID
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
}

func (h *PersonHandler) GetPeopleToken(
	w http.ResponseWriter,
	r *http.Request,
) {

	limit := 20

	limitParam := r.URL.Query().Get("limit")

	if limitParam != "" {

		parsedLimit, err := strconv.Atoi(limitParam)

		if err != nil ||
			parsedLimit < 1 ||
			parsedLimit > 100 {

			http.Error(
				w,
				"Limit must be between 1 and 100",
				http.StatusBadRequest,
			)

			return
		}

		limit = parsedLimit
	}

	paginationToken := r.URL.Query().Get("token")

	var playerID string

	/*
		First request:

		/people/token?limit=10

		No token exists.
	*/

	if paginationToken != "" {

		var err error

		playerID, err = token.Decode(
			paginationToken,
		)

		if err != nil {

			http.Error(
				w,
				"Invalid pagination token",
				http.StatusBadRequest,
			)

			return
		}
	}

	/*
		Ask database for one extra record.

		limit = 10

		database returns 11.

		10 → actual response
		1  → tells us another page exists
	*/

	people, err := h.Service.GetPeopleToken(
		r.Context(),
		playerID,
		int32(limit+1),
	)

	if err != nil {

		log.Println(
			"GetPeopleToken error:",
			err,
		)

		http.Error(
			w,
			"Failed to fetch people",
			http.StatusInternalServerError,
		)

		return
	}

	hasNext := len(people) > limit

	if hasNext {
		people = people[:limit]
	}

	response := TokenPaginatedResponse{
		Data:    people,
		HasNext: hasNext,
	}

	if hasNext {

		lastPerson := people[len(people)-1]

		nextToken, err := token.Create(
			lastPerson.PlayerID,
		)

		if err != nil {

			log.Println(
				"Create pagination token error:",
				err,
			)

			http.Error(
				w,
				"Failed to create pagination token",
				http.StatusInternalServerError,
			)

			return
		}

		response.NextToken = nextToken
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
}

/*
=========================================================
UPDATE PERSON
=========================================================

PUT /people/update?playerID=xxxxx

playerID comes from the URL query parameter and
cannot be changed.

All other fields can be updated.
*/

func (h *PersonHandler) UpdatePerson(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Only PUT is allowed
	if r.Method != http.MethodPut {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	// Get playerID from query parameter
	playerID := r.URL.Query().Get("playerID")

	if playerID == "" {
		http.Error(
			w,
			"Player ID is required",
			http.StatusBadRequest,
		)
		return
	}

	// Request body
	var request struct {

		BirthYear    *int16  `json:"birthYear"`
		BirthMonth   *int16  `json:"birthMonth"`
		BirthDay     *int16  `json:"birthDay"`
		BirthCountry *string `json:"birthCountry"`
		BirthState   *string `json:"birthState"`
		BirthCity    *string `json:"birthCity"`

		DeathYear    *int16  `json:"deathYear"`
		DeathMonth   *int16  `json:"deathMonth"`
		DeathDay     *int16  `json:"deathDay"`
		DeathCountry *string `json:"deathCountry"`
		DeathState   *string `json:"deathState"`
		DeathCity    *string `json:"deathCity"`

		NameFirst *string `json:"nameFirst"`
		NameLast  *string `json:"nameLast"`
		NameGiven *string `json:"nameGiven"`

		Weight *int16 `json:"weight"`
		Height *int16 `json:"height"`

		Bats   *string `json:"bats"`
		Throws *string `json:"throws"`

		Debut     *string `json:"debut"`
		FinalGame *string `json:"finalGame"`

		RetroID *string `json:"retroID"`
		BbrefID *string `json:"bbrefID"`
	}

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(
			w,
			"Invalid JSON body",
			http.StatusBadRequest,
		)
		return
	}

	// Convert request to SQLC parameters
	params := generated.UpdatePersonParams{

		// playerID is NOT taken from JSON.
		// It comes from the URL.
		Playerid: playerID,

		Birthyear: sql.NullInt16{
			Int16: valueOrZero(request.BirthYear),
			Valid: request.BirthYear != nil,
		},

		Birthmonth: sql.NullInt16{
			Int16: valueOrZero(request.BirthMonth),
			Valid: request.BirthMonth != nil,
		},

		Birthday: sql.NullInt16{
			Int16: valueOrZero(request.BirthDay),
			Valid: request.BirthDay != nil,
		},

		Birthcountry: nullString(request.BirthCountry),
		Birthstate:   nullString(request.BirthState),
		Birthcity:    nullString(request.BirthCity),

		Deathyear: sql.NullInt16{
			Int16: valueOrZero(request.DeathYear),
			Valid: request.DeathYear != nil,
		},

		Deathmonth: sql.NullInt16{
			Int16: valueOrZero(request.DeathMonth),
			Valid: request.DeathMonth != nil,
		},

		Deathday: sql.NullInt16{
			Int16: valueOrZero(request.DeathDay),
			Valid: request.DeathDay != nil,
		},

		Deathcountry: nullString(request.DeathCountry),
		Deathstate:   nullString(request.DeathState),
		Deathcity:    nullString(request.DeathCity),

		Namefirst: nullString(request.NameFirst),
		Namelast:  nullString(request.NameLast),
		Namegiven: nullString(request.NameGiven),

		Weight: sql.NullInt16{
			Int16: valueOrZero(request.Weight),
			Valid: request.Weight != nil,
		},

		Height: sql.NullInt16{
			Int16: valueOrZero(request.Height),
			Valid: request.Height != nil,
		},

		Bats:   nullString(request.Bats),
		Throws: nullString(request.Throws),

		Debut:     nullTime(request.Debut),
		Finalgame: nullTime(request.FinalGame),

		Retroid: nullString(request.RetroID),
		Bbrefid: nullString(request.BbrefID),
	}

	// Call service
	err = h.Service.UpdatePerson(
		r.Context(),
		params,
	)

	if err != nil {

		log.Println(
			"UpdatePerson error:",
			err,
		)

		http.Error(
			w,
			"Failed to update person",
			http.StatusInternalServerError,
		)

		return
	}

	// Success response
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Person updated successfully",
		"playerID": playerID,
	})
}

func valueOrZero(value *int16) int16 {

	if value == nil {
		return 0
	}

	return *value
}

func nullString(value *string) sql.NullString {

	if value == nil {
		return sql.NullString{}
	}

	return sql.NullString{
		String: *value,
		Valid:  true,
	}
}

func nullTime(value *string) sql.NullTime {

	if value == nil || *value == "" {
		return sql.NullTime{}
	}

	parsed, err := time.Parse(
		"2006-01-02",
		*value,
	)

	if err != nil {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  parsed,
		Valid: true,
	}
}

func (h *PersonHandler) GetAllPeople(
    w http.ResponseWriter,
    r *http.Request,
) {

    page := 1
    limit := 20

    pageParam := r.URL.Query().Get("page")
    limitParam := r.URL.Query().Get("limit")

    sortBy := r.URL.Query().Get("sortBy")
    sortOrder := r.URL.Query().Get("sortOrder")

    var err error

    if pageParam != "" {
        page, err = strconv.Atoi(pageParam)

        if err != nil || page < 1 {
            http.Error(
                w,
                "Invalid page",
                http.StatusBadRequest,
            )
            return
        }
    }

    if limitParam != "" {
        limit, err = strconv.Atoi(limitParam)

        if err != nil || limit < 1 || limit > 100 {
            http.Error(
                w,
                "Limit must be between 1 and 100",
                http.StatusBadRequest,
            )
            return
        }
    }

    // Default sorting
    if sortBy == "" {
        sortBy = "playerID"
    }

    if sortOrder == "" {
        sortOrder = "asc"
    }

    // Validate sort field
    validSortFields := map[string]bool{
        "nameFirst":  true,
        "birthYear":  true,
        "height":     true,
        "playerID":   true,
    }

    if !validSortFields[sortBy] {
        http.Error(
            w,
            "Invalid sort field",
            http.StatusBadRequest,
        )
        return
    }

    // Validate sort order
    if sortOrder != "asc" && sortOrder != "desc" {
        http.Error(
            w,
            "Sort order must be asc or desc",
            http.StatusBadRequest,
        )
        return
    }

    offset := (page - 1) * limit

    people, err := h.Service.GetPeoplePaginatedSorted(
        r.Context(),
        int32(limit),
        int32(offset),
        sortBy,
        sortOrder,
    )

    if err != nil {
        log.Println(
            "GetPeoplePaginatedSorted error:",
            err,
        )

        http.Error(
            w,
            "Failed to fetch people",
            http.StatusInternalServerError,
        )
        return
    }

    total, err := h.Service.CountPeople(r.Context())

    if err != nil {
        log.Println(
            "CountPeople error:",
            err,
        )

        http.Error(
            w,
            "Failed to count people",
            http.StatusInternalServerError,
        )
        return
    }

    totalPages := int(
        (total + int64(limit) - 1) / int64(limit),
    )

    if page > totalPages && totalPages > 0 {
        http.Error(
            w,
            "Page not found",
            http.StatusNotFound,
        )
        return
    }

    response := PaginatedResponse{
        Data:       people,
        Page:       page,
        Limit:      limit,
        Total:      total,
        TotalPages: totalPages,
    }

    w.Header().Set(
        "Content-Type",
        "application/json",
    )

    json.NewEncoder(w).Encode(response)
}