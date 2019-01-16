package server

import (
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
