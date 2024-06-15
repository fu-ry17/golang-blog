package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is required")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}

func ConvertBlogIdToInt(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("blog id is required")
	}

	intId, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("blog id is invalid")
	}

	return intId, nil
}
