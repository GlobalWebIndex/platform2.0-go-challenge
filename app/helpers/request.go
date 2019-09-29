package helpers

import (
	"errors"
	"net/http"
)

func ExportParam(r *http.Request, paramName string, defaultValue string) (string, error) {
	returnValue := defaultValue
	query_values := r.URL.Query() // Returns a url query_values, which is a map[string][]string
	paramValues, ok := query_values[paramName]
	if ok {
		if len(paramValues) >= 1 {
			returnValue = paramValues[0] // The first `?paramName=???`
		}
	}
	if !ok {
		return returnValue, errors.New("Parameter not found")
	}

	return returnValue, nil
}
