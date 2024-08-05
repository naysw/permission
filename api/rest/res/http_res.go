package res

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Status  int         `json:"status,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type responseOption func(*response)

func WithMessage(message string) responseOption {
	return func(r *response) {
		r.Message = message
	}
}

func WithData(data interface{}) responseOption {
	return func(r *response) {
		r.Data = data
	}
}

func WithStatusCode(statusCode int) responseOption {
	return func(r *response) {
		r.Status = statusCode
	}
}

func WithMeta(meta interface{}) responseOption {
	return func(r *response) {
		r.Meta = meta
	}
}

func WithErrors(errors interface{}) responseOption {
	return func(r *response) {
		r.Errors = errors
	}
}

func NewHttpErr(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError
	msg := http.StatusText(http.StatusInternalServerError)

	writeJSON(
		w,
		statusCode,
		response{
			Message: msg,
			Status:  statusCode,
		},
	)
}

func NewBadRequest(
	w http.ResponseWriter,
	opts ...responseOption) {

	dd := response{
		Message: http.StatusText(http.StatusBadRequest),
		Status:  http.StatusBadRequest,
	}

	for _, opt := range opts {
		opt(&dd)
	}

	writeJSON(w, dd.Status, dd)
}

func NewNotFound(w http.ResponseWriter, opts ...responseOption) {
	dd := response{
		Message: http.StatusText(http.StatusNotFound),
		Status:  http.StatusNotFound,
	}

	for _, opt := range opts {
		opt(&dd)
	}

	writeJSON(w, dd.Status, dd)
}

func NewUnauthorized(w http.ResponseWriter, opts ...responseOption) {
	dd := response{
		Message: http.StatusText(http.StatusUnauthorized),
		Status:  http.StatusUnauthorized,
	}

	for _, opt := range opts {
		opt(&dd)
	}

	writeJSON(w, dd.Status, dd)
}

func NewTooManyRequests(w http.ResponseWriter, opts ...responseOption) {
	dd := response{
		Message: http.StatusText(http.StatusTooManyRequests),
		Status:  http.StatusTooManyRequests,
	}
	for _, opt := range opts {
		opt(&dd)
	}

	writeJSON(w, dd.Status, dd)
}

func NewUnprocessableEntity(w http.ResponseWriter, opts ...responseOption) {
	dd := response{
		Message: "validation failed",
		Status:  http.StatusUnprocessableEntity,
	}
	for _, opt := range opts {
		opt(&dd)
	}

	writeJSON(
		w,
		http.StatusUnprocessableEntity,
		dd,
	)
}

func NewOK(
	w http.ResponseWriter,
	opts ...responseOption) {

	dd := response{
		Message: http.StatusText(http.StatusOK),
		Status:  http.StatusOK,
	}

	for _, opt := range opts {
		opt(&dd)
	}

	writeJSON(w, dd.Status, dd)
}

func NewCreated(
	w http.ResponseWriter,
	opts ...responseOption) {

	dd := response{
		Message: http.StatusText(http.StatusCreated),
		Status:  http.StatusCreated,
	}

	for _, opt := range opts {
		opt(&dd)
	}

	writeJSON(w, dd.Status, dd)
}

func writeJSON(
	w http.ResponseWriter,
	statusCode int,
	data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
