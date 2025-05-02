package handlers_builder

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

func handlerResultFunc(ctx context.Context, w http.ResponseWriter, r *http.Request, result any) {
	bs, err := json.Marshal(makeResultBody(result))
	if err != nil {
		slog.Error("error marshalling json", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(bs)
}
