package signin

import (
	"context"
	"errors"
	"net/http"

	ergo "github.com/nktknshn/go-ergo-handler"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

type HandlerSignIn struct {
	authSignin usecase
}

type usecase interface {
	Handle(ctx context.Context, in auth_signin.In) (auth_signin.Out, error)
	GetName() string
}

// @Summary      Sign in
// @ID           signIn
// @Description  Sign in
// @Tags         signin
// @Accept       json
// @Produce      json
// @Param        payload   body      requestBody  true  "Payload"
// @Success      200  {object}  handlers_builder.Result[responseBody]
// @Failure      400  {object}  handlers_builder.Error
// @Failure      401  {object}  handlers_builder.Error
// @Failure      500  {object}  handlers_builder.Error
// @Router       /api/v1/signin [post]
func New(authSignin usecase) *HandlerSignIn {

	if authSignin == nil {
		panic("authSignin is nil")
	}

	return &HandlerSignIn{authSignin: authSignin}
}

func (h *HandlerSignIn) GetHandler() http.Handler {
	return makeHandlerSignIn(h.authSignin)
}

type requestBody struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"admin1234"`
}

func (p requestBody) GetIn() (auth_signin.In, error) {
	return auth_signin.NewInFromValues(
		p.Username,
		p.Password,
	)
}

type responseBody struct {
	Token string `json:"token"`
}

func makeHandlerSignIn(u usecase) http.Handler {
	var (
		b       = handlers_builder.NewPublic()
		payload = ergo.PayloadAttach[requestBody](b)
	)

	return b.BuildHandlerWrapped(func(_ http.ResponseWriter, r *http.Request) (any, error) {
		pl := payload.Get(r)
		in, err := pl.GetIn()

		if err != nil {
			return nil, ergo.NewError(http.StatusBadRequest, err)
		}

		out, err := u.Handle(r.Context(), in)

		if errors.Is(err, domainAuth.ErrInvalidAuthUserPassword) {
			return nil, ergo.NewError(http.StatusUnauthorized, err)
		}

		if err != nil {
			return nil, err
		}

		return responseBody{Token: out.Token.String()}, nil
	})
}

func (h *HandlerSignIn) GetName() string {
	return use_cases.NameAuthSignin
}
