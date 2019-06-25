package main

import (
	"encoding/json"
	"net/http"

	snowflake "github.com/piaohu/snowflake-golang"
)

func init() {
	snowflake.DefaultNode()
}

func handler(w http.ResponseWriter, r *http.Request) {
	id := make(map[string]interface{}, 0)
	id["id"] = snowflake.Generate()
	body, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
