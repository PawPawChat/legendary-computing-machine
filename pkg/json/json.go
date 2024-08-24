package json

import (
	"encoding/json"
	"io"
	"log"
)

func MustMarshal(v any) []byte {
	m, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func Decode(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}
