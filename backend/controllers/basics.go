package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ivan-ca97/rush/backend/custom_errors"
	ce "github.com/ivan-ca97/rush/backend/custom_errors"
)

func (rc *RushControllers) GetPaginationParameters(request *http.Request) (pageNumber int, pageSize int, err error) {
	pageNumberString := chi.URLParam(request, "page_number")
	pageSizeString := chi.URLParam(request, "page_size")

	if pageNumberString == "" {
		pageNumberString = "0"
	}
	if pageSizeString == "" {
		pageSizeString = "0"
	}

	pageNumber, err1 := strconv.Atoi(pageNumberString)
	pageSize, err2 := strconv.Atoi(pageSizeString)
	if err1 != nil || err2 != nil {
		err := &ce.RequestError{Message: "Page number or size invalid"}
		return 0, 0, err
	}

	return pageNumber, pageSize, nil
}

func (rc *RushControllers) EncodeResponse(w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return &ce.InternalServerError{Message: "Server error"}
	}

	return nil
}

func (rc *RushControllers) HandleError(w http.ResponseWriter, err error) {
	var expectedErr custom_errors.ExpectedError

	if errors.As(err, &expectedErr) {
		w.WriteHeader(expectedErr.StatusCode())
		json.NewEncoder(w).Encode(expectedErr)
		return
	}

	log.Printf("Unexpected error: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error"})
}
