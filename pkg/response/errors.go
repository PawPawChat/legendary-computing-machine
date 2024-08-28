package response

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"google.golang.org/grpc/status"
)

func WriteProtoError(w http.ResponseWriter, err error) {
	type convert struct {
		Message string
		Details []any
	}

	status := status.Convert(err)

	pberr := convert{
		Message: status.Message(),
		Details: status.Details(),
	}

	slog.With("error", err).Error("")
	Json().CodeFromProto(status.Code()).Body(map[string]any{"error": pberr}).MustWrite(w)
}

func WriteMissingFieldsError(w http.ResponseWriter, fields []string) {
	type convert struct {
		Message string
		Fields  []string
	}

	conv := convert{
		Message: "missing requirements fields",
		Fields:  fields,
	}

	slog.Error(conv.Message)
	Json().BadRequest().Body(map[string]any{"error": conv}).MustWrite(w)
}

func WriteParseBodyError(w http.ResponseWriter, err error) {
	type convert struct {
		Message string
		Field   string
		Reason  string
	}

	conv := convert{
		Message: "failed to parse the request body",
	}

	switch jserr := err.(type) {
	case *json.SyntaxError:
		conv.Reason = fmt.Sprintf("syntax error at byte offset %d", jserr.Offset)
	case *json.UnmarshalTypeError:
		conv.Field = jserr.Field
		conv.Reason = fmt.Sprintf("expected type %s but got %s", jserr.Type, jserr.Value)
	default:
		conv.Reason = err.Error()
	}

	slog.Error(conv.Message, "error", err)
	Json().Body(map[string]any{"error": conv}).BadRequest().MustWrite(w)
}
