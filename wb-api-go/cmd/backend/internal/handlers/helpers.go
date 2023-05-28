package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error  `json:"-"` // low-level runtime error
	HTTPStatusCode int    `json:"http_status_code,omitempty"`
	StackTrace     string `json:"stack_trace,omitempty"`
	StatusText     string `json:"status,omitempty"` // user-level status message
	AppCode        int64  `json:"code,omitempty"`   // application-specific error code
	ErrorText      string `json:"error,omitempty"`  // application-level error message, for debugging
}

type ApiResponse struct {
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"status_code,omitempty"`
	StatusText string      `json:"status_text,omitempty"`
}

func (e *ErrResponse) Error() string {
	return e.Err.Error()
}

func (h *Handler) text(w http.ResponseWriter, r *http.Request, data []byte) error {
	w.Header().Set("Content-Type", "text/plain")
	render.Status(r, 200)
	w.Write(data)
	return nil
}

func (h *Handler) success(w http.ResponseWriter, r *http.Request, v interface{}) error {
	render.Status(r, 200)
	render.JSON(w, r, ApiResponse{
		Data:       v,
		StatusCode: 200,
		StatusText: "success",
	})
	return nil
}

func (h *Handler) status(w http.ResponseWriter, r *http.Request, status int, v interface{}) error {
	render.Status(r, status)
	render.JSON(w, r, v)
	return nil
}

func (h *Handler) error(w http.ResponseWriter, r *http.Request, status int, err error) error {
	response := ErrResponse{
		Err:            err,
		HTTPStatusCode: status,
		StatusText:     http.StatusText(status),
		ErrorText:      err.Error(),
	}

	response.Render(w, r)
	return &response
}

// Render  renders rendery things to render wtse-1
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	render.Status(r, e.HTTPStatusCode)
	render.JSON(w, r, e)
	return nil
}
