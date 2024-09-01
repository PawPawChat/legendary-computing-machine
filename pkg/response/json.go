package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// jsonBuilder is a utility to build JSON responses for HTTP handlers. It handles
// structuring the body, setting response codes, and writing the response to the client
type jsonBuilder struct {
	body []byte
	code int
	howl error
}

func Json() *jsonBuilder {
	return &jsonBuilder{}
}

// Body adds objects to the json response body
func (b *jsonBuilder) Body(object any) *jsonBuilder {
	b.body, b.howl = json.Marshal(object)
	return b
}

// MustWrite writes the JSON response to the provided http.ResponseWriter
// It will log a fatal error if an error occurs during the writing process
func (b *jsonBuilder) MustWrite(w http.ResponseWriter) {
	b.setHeader(w)
	if _, err := w.Write(b.body); err != nil {
		log.Fatal(err)
	}
}

// Write writes the JSON response to the provided http.ResponseWriter
// It returns an error if any issues occur during the response construction or writing proces
func (b *jsonBuilder) Write(w http.ResponseWriter) error {
	if b.howl != nil {
		return b.howl
	}
	b.setHeader(w)
	if _, err := w.Write(b.body); err != nil {
		return err
	}
	return nil
}

// Code sets the HTTP header code for the response
func (b *jsonBuilder) setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(b.code)
}
