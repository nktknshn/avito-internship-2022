package handlers_params

import (
	"github.com/nktknshn/avito-internship-2022/internal/domain"
	ergo "github.com/nktknshn/go-ergo-handler"
)

var (
	RouterParamUserID = ergo.RouterParam("user_id", ergo.IgnoreContext(domain.NewUserIDFromString))
)
