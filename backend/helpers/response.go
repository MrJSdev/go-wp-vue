package helpers

import (
	"encoding/json"
	"net/http"
)

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	ResponseWithJSON(w, code, map[string]string{"error": msg})
}

func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Access-Control-Allow-Origin", "localhost:3000")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
