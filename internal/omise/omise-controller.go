package omise

import (
	"git.akyoto.dev/go/web"
)

type OmiseController struct {
	Callback func(ctx web.Context) error
}

func New() *OmiseController {
	return &OmiseController{}
}

func Callback(ctx web.Context) error {
	return ctx.String("Success")
}
