package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	var res map[string]interface{}
	if code >= http.StatusOK && code < http.StatusMultipleChoices {
		// Códigos 2xx (éxito)
		if data == nil {
			res = map[string]interface{}{
				"status":  "success",
				"message": message,
			}
		} else {
			res = map[string]interface{}{
				"status":  "success",
				"message": message,
				"data":    data,
			}
		}
	} else if code >= http.StatusBadRequest && code < http.StatusInternalServerError {
		res = map[string]interface{}{
			"status":  "fail",
			"message": message,
		}
	} else if code >= http.StatusInternalServerError {
		res = map[string]interface{}{
			"status":  "error",
			"message": message,
		}
	}

	json.NewEncoder(w).Encode(res)
}
