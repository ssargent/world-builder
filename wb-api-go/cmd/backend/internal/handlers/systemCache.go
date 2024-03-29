package handlers

import (
	"context"
	"net/http"
)

func (h *Handler) systemCache(_ context.Context, w http.ResponseWriter, r *http.Request) error {
	cached := h.cache.Items()
	cacheList := make([]string, len(cached))

	idx := 0

	for k := range cached {
		cacheList[idx] = k
		idx++
	}

	return h.success(w, r, cacheList)
}
