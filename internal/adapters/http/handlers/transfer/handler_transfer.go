package transfer

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
	"github.com/nktknshn/avito-internship-2022/internal/use_cases/transfer"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type TransferHandler struct {
	auth    handlers_auth.TokenValidator
	useCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in transfer.In) error
}

func NewTransferHandler(auth handlers_auth.TokenValidator, useCase useCase) *TransferHandler {

	if auth == nil {
		panic("auth is nil")
	}

	if useCase == nil {
		panic("useCase is nil")
	}

	return &TransferHandler{auth: auth, useCase: useCase}
}

func (h *TransferHandler) GetHandler() http.Handler {
	return makeTransferHandler(h.auth, h.useCase)
}

type payloadType struct {
	From   int64 `json:"from"`
	To     int64 `json:"to"`
	Amount int64 `json:"amount"`
}

func (p payloadType) GetIn() (transfer.In, error) {
	return transfer.NewInFromValues(p.From, p.To, p.Amount)
}

func makeTransferHandler(auth handlers_auth.TokenValidator, u useCase) http.Handler {
	var (
		b, _ = handlers_builder.NewWithAuth(auth, []domain.AuthUserRole{
			domain.AuthUserRoleAdmin,
			domain.AuthUserRoleAccount,
		})
		payload = ergo.PayloadAttach[payloadType](b)
	)

	return b.BuildHandlerWrapped(func(w http.ResponseWriter, r *http.Request) (any, error) {
		pl := payload.Get(r)
		in, err := pl.GetIn()
		if err != nil {
			return nil, ergo.NewError(http.StatusBadRequest, err)
		}
		err = u.Handle(r.Context(), in)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
}
