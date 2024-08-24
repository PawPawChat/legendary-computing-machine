package response

import (
	"log"
	"net/http"

	"github.com/pawpawchat/core/pkg/json"
)

type ResponseBuilder struct {
	writer      http.ResponseWriter
	code        int
	contentType string
	body        []byte
}

func Json(w http.ResponseWriter) *ResponseBuilder {
	return &ResponseBuilder{
		writer:      w,
		body:        make([]byte, 0),
		contentType: "application/json",
	}
}

func (b *ResponseBuilder) Code(code int) *ResponseBuilder {
	b.code = code
	return b
}

func (b *ResponseBuilder) Args(args ...string) *ResponseBuilder {
	m := make(map[string]interface{})
	for idx := range len(args) - 1 {
		m[args[idx]] = args[idx+1]
	}
	b.body = json.MustMarshal(m)
	return b
}

func (b *ResponseBuilder) Map(bodyMap map[string]interface{}) *ResponseBuilder {
	b.body = json.MustMarshal(bodyMap)
	return b
}

func (b *ResponseBuilder) Body(bytes []byte) *ResponseBuilder {
	b.body = bytes
	return b
}

func (b *ResponseBuilder) MustSend() {
	b.writer.Header().Set("Content-Type", b.contentType)
	b.writer.WriteHeader(b.code)
	if _, err := b.writer.Write(b.body); err != nil {
		log.Fatal(err)
	}
}
