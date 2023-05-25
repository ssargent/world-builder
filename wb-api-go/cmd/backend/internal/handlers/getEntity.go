package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (h *Handler) getEntity(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	idstr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		return h.error(w, r, 400, err)
	}

	ent, err := h.entity.FindByID(ctx, id)
	if err != nil {
		return h.error(w, r, 500, err)
	}

	return h.success(w, r, ent)
}
