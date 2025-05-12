package handlers_builder

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

func handlerResultFunc(_ context.Context, w http.ResponseWriter, _ *http.Request, result any) {
	bs, err := json.Marshal(makeResultBody(result))
	if err != nil {
		// TODO: fix this
		//nolint:sloglint // позже придумать, как сделать. Может через контекст?
		slog.Error("error marshalling json", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(bs)
	if err != nil {
		//nolint:sloglint // позже придумать, как сделать. Может через контекст?
		slog.Error("error writing response", "error", err)
	}
}
