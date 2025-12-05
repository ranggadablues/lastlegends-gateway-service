package helper

import (
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code"`
	Meta    Meta        `json:"meta"`
}

type Meta struct {
	Timestamp string `json:"timestamp"`
}

var protoMarshaler = protojson.MarshalOptions{
	EmitUnpopulated: true,
	UseProtoNames:   false, // camelCase for UI
}

// Send success JSON (supports protobuf or normal structs)
func WriteSuccessHeader(w http.ResponseWriter, data interface{}) {
	var anyData interface{}

	// If data is protobuf, marshal to JSON first
	if msg, ok := data.(proto.Message); ok {
		b, err := protoMarshaler.Marshal(msg)
		if err == nil {
			_ = json.Unmarshal(b, &anyData)
		} else {
			gRPCError(w, http.StatusInternalServerError, "failed to encode protobuf data")
			return
		}
	} else {
		anyData = data
	}

	resp := APIResponse{
		Success: true,
		Data:    anyData,
		Code:    http.StatusOK,
		Meta:    Meta{Timestamp: time.Now().UTC().Format(time.RFC3339)},
	}

	writeJSON(w, http.StatusOK, resp)
}

// --- gRPC → HTTP error converter ---
func gRPCError(w http.ResponseWriter, code int, message string) {
	resp := APIResponse{
		Success: false,
		Message: message,
		Code:    code,
		Meta:    Meta{Timestamp: time.Now().UTC().Format(time.RFC3339)},
	}
	writeJSON(w, code, resp)
}

// Send error JSON
func WriteErrorHeader(w http.ResponseWriter, err error) {
	st, ok := status.FromError(err)
	if !ok {
		gRPCError(w, http.StatusInternalServerError, "unknown internal error")
		return
	}

	code := grpcToHTTP(st.Code())
	gRPCError(w, code, st.Message())
}

// --- Helper to map gRPC → HTTP status codes ---
func grpcToHTTP(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

// Internal helper
func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
