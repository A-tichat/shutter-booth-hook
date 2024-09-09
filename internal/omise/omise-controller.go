package omise

import (
	"git.akyoto.dev/go/web"
)

type OmiseController struct {
}

func New() *OmiseController {
	return &OmiseController{}
}

func (ctr OmiseController) Callback(ctx web.Context) error {
	return ctx.String("Success")
}
