package server

import (
	"akovalyov/chlorine/apierror"
	"encoding/json"
	"net/http"
)

// JSONResponseWriter structure adds methods to write JSON to the ResponseWriter.
type JSONResponseWriter struct {
	http.ResponseWriter
}

// WriteJSON writes sequence of bytes to ResponseWriter and adds Content-Type header as "application/json".
func (w JSONResponseWriter) WriteJSON(data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

// WriteJSONObject marshals structure and writes it to the ResponseWriter with Content-Type set to "application/json".
func (w JSONResponseWriter) WriteJSONObject(object interface{}) error {
	marshaledObject, err := json.Marshal(object)
	if err != nil {
		return err
	}
	w.WriteJSON(marshaledObject)
	return nil
}

// Error writes to the ResponseWriter APIError in JSON serialized form.
func (w JSONResponseWriter) Error(apiError apierror.APIError, httpCode int) {
	w.Header().Add("Content-Type", "application/json")
	errMsg, err := json.Marshal(apiError)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(httpCode)
	w.Write(errMsg)

}
