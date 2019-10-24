package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status string, messsage string) (map[string]interface{})  {
	return map[string]interface{}{"status": status, "message":messsage}
}

func Respond(w http.ResponseWriter, data map[string]interface{})  {
	w.Header().Add("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(data)
}