package response

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"time"

	"google.golang.org/grpc/status"
)

func WriteProtoError(w http.ResponseWriter, err error) {
	type convert struct {
		Message string `json:"message"`
		Details []any  `json:"details"`
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
		Message string   `json:"message"`
		Fields  []string `json:"fields"`
	}

	conv := convert{
		Message: "missing requirements fields",
		Fields:  fields,
	}

	slog.Error(conv.Message)
	Json().BadRequest().Body(map[string]any{"error": conv}).MustWrite(w)
}

func WriteParseBodyError(w http.ResponseWriter, err error) {
	var convert any

	switch err := err.(type) {
	case *json.SyntaxError:
		convert = struct {
			Message string `json:"message"`
			Offset  string `json:"offset"`
		}{
			"failed to parse request body, syntax error",
			fmt.Sprintf("at byte %d", err.Offset),
		}
	case *json.UnmarshalTypeError:
		convert = struct {
			Message string `json:"message"`
			Field   string `json:"field"`
			Details string `json:"details"`
		}{
			"failed to parse request body",
			err.Field,
			fmt.Sprintf("expected type %s but got %s", err.Type, err.Value),
		}
	case *time.ParseError:
		convert = struct {
			Message  string `json:"message"`
			Expected string `json:"expected"`
			Received string `json:"received"`
		}{
			"failed to parse request body, time parsing",
			err.Layout,
			err.ValueElem,
		}
	default:
		slog.Debug("unexpected error", "type", reflect.TypeOf(err), "msg", err.Error())
		convert = err
	}

	slog.Error("", "msg", convert)
	Json().Body(map[string]any{"error": convert}).BadRequest().MustWrite(w)
}
